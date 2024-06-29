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
	Storer
	IDGenerator
	IoConnector
	EventHandler
	Txer

	Logger *slog.Logger

	ctx context.Context

	liquid *liquid

	bg Backgrounder

	config *Config

	isRunning bool
}

func NewEngine(name string, ops ...Option) *Engine {
	config := &Config{
		EventHandler: eh{},
		Context:      context.Background(),
		IoConnector:  nopIo{},
		Logger:       slog.Default(),
	}

	for _, op := range ops {
		op(config)
	}

	e := &Engine{
		Name:        name,
		Evaluator:   NewExprEval(),
		IDGenerator:       uid{},
		IoConnector: config.IoConnector,

		TemplateGetter: config.templates,
		Storer:          config.Store,
		Txer:           config.Txer,
		EventHandler:   config.EventHandler,

		Logger: config.Logger.With(slog.String("engine", name)),

		ctx: config.Context,

		config: config,
	}

	e.bg = bg{engine: e}

	e.liquid = &liquid{
		engine: e,

		actCh: make(chan Activity, 1),
		ech:   make(chan Activity, 1),

		logger: config.Logger.With(slog.String("driver", "liquid")),
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
	return e.bg.Background(cmd)
}
