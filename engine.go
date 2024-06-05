package loong

import (
	"context"
	"errors"
)

type Engine struct {
	Name string

	Evaluator
	Templates
	Store
	IDGen
	IoConnector

	ctx context.Context

	driver Driver

	config *Config

	isRunning bool
}

func NewEngine(name string, ops ...Option) *Engine {
	config := &Config{
		queueSize: 4,
		eh:        emptyEh,
		ctx:       context.Background(),
		connector: emptyConnect,
	}

	for _, op := range ops {
		op(config)
	}

	return &Engine{
		Name:        name,
		Evaluator:   NewExprEval(),
		IDGen:       uid{},
		IoConnector: config.connector,

		Templates: config.templates,
		Store:     config.store,

		ctx: config.ctx,

		driver: newLiquid(config.ctx, config.eh, config.queueSize),

		config: config,
	}
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
