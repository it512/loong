package loong

import (
	"context"
	"errors"
	"log/slog"
)

type Engine struct {
	Name string

	Evaluator
	Templates
	Store
	IDGen
	IoConnector
	EventHandler

	Logger *slog.Logger

	ctx context.Context

	driver Driver

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

		Templates:    config.templates,
		Store:        config.store,
		EventHandler: config.eh,

		Logger: config.logger.With(slog.String("engine", name)),

		ctx: config.ctx,

		config: config,
	}

	e.driver = &liquid{
		loop: make(chan Activity, 1),
		c:    config.ctx,

		ech: make(chan Activity, 1),
		eh:  config.eh,

		logger: config.logger.With(slog.String("driver", "liquid")),

		size: 4,
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

	if err = e.driver.Run(); err != nil {
		return
	}

	logo()

	e.isRunning = true
	return
}

func (e *Engine) CommitCmd(ctx context.Context, cmd ActivityCmd) error {
	if err := cmd.Init(ctx, e); err != nil {
		return err
	}
	return e.Emit(cmd)
}

/*
func (e *Engine) RunActivityCmd(ctx context.Context, cmd ActivityCmd) error {
	if err := cmd.Init(ctx, e); err != nil {
		return err
	}
	return e.Emit(cmd)
}
*/

func (e *Engine) RunCmd(ctx context.Context, cmd Cmd) error {
	return e.CommitCmd(ctx, &bgCmd{Cmd: cmd})
}

func (e *Engine) Emit(ops ...Activity) error {
	if !e.isRunning {
		return errors.New("引擎未运行")
	}
	return e.driver.Emit(ops...)
}

type bgCmd struct{ Cmd }

func (bgCmd) Emit(ctx context.Context, emt Emitter) error { return nil }
func (bgCmd) Type() ActivityType                          { return backgroundCmd }
