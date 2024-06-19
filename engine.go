package loong

import (
	"context"
	"errors"
	"log/slog"
)

type Engine struct {
	Name string

	Evaluator
	TemplateGetter
	Store
	IDGen
	IoConnector
	EventHandler
	Txer

	Logger *slog.Logger

	ctx context.Context

	liquid *liquid

	config *Config

	isRunning bool
}

func NewEngine(name string, ops ...Option) *Engine {
	config := &Config{
		queueSize: 4,
		eh:        eh{},
		ctx:       context.Background(),
		connector: nopIo{},
		logger:    slog.Default(),
	}

	for _, op := range ops {
		op(config)
	}

	e := &Engine{
		Name:        name,
		Evaluator:   NewExprEval(),
		IDGen:       uid{},
		IoConnector: config.connector,

		TemplateGetter: config.templates,
		Store:          config.store,
		EventHandler:   config.eh,

		Logger: config.logger.With(slog.String("engine", name)),

		ctx: config.ctx,

		config: config,
	}

	e.liquid = &liquid{
		engine: e,

		cmdCh: make(chan Cmd, 1),
		actCh: make(chan Activity, 1),
		ech:   make(chan Activity, 1),

		logger: config.logger.With(slog.String("driver", "liquid")),
		size:   4,
	}

	return e
}

func (e *Engine) init() error {
	return nil
}

func (e *Engine) Run() (err error) {
	if e.isRunning {
		return
	}

	if err = e.init(); err != nil {
		return
	}

	if err = e.liquid.run(); err != nil {
		return
	}

	logo()

	e.isRunning = true
	return
}

func (e *Engine) RunActivityCmd(ctx context.Context, activity ActivityCmd) error {
	if !e.isRunning {
		return errors.New("引擎未运行")
	}
	if err := activity.Bind(ctx, e); err != nil {
		return err
	}
	return e.liquid.Emit(activity)
}

func (e *Engine) BackgroundCmd(ctx context.Context, cmd Cmd) error {
	if !e.isRunning {
		return errors.New("引擎未运行")
	}
	if err := cmd.Bind(ctx, e); err != nil {
		return err
	}
	return e.liquid.Background(cmd)
}
