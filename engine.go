package loong

import (
	"context"
	"errors"
)

type Engine struct {
	name string

	Evaluator
	Templates
	Store
	IDGen
	IoConnector
	EventHandler

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
	}

	for _, op := range ops {
		op(config)
	}

	e := &Engine{
		name:        name,
		Evaluator:   NewExprEval(),
		IDGen:       uid{},
		IoConnector: config.connector,

		Templates:    config.templates,
		Store:        config.store,
		EventHandler: config.eh,

		ctx: config.ctx,

		config: config,
	}
	e.driver = newLiquid(config.ctx, e, config.queueSize)
	return e
}

func (e *Engine) Name() string {
	return e.name
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

	Logo()

	e.isRunning = true
	return
}

func (e *Engine) CommitCmd(ctx context.Context, cmd ActivityCmd) error {
	if err := cmd.Init(ctx, e); err != nil {
		return err
	}
	return e.Emit(cmd)
}

func (e *Engine) RunCmd(ctx context.Context, cmd Cmd) error {
	if err := cmd.Init(ctx, e); err != nil {
		return err
	}
	return cmd.Do(e.ctx)
}

func (e *Engine) Emit(ops ...Activity) error {
	if !e.isRunning {
		return errors.New("引擎未运行")
	}
	return e.driver.Emit(ops...)
}
