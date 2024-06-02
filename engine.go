package loong

import (
	"context"
	"errors"
	"fmt"
)

type StartProcCmd struct {
	ProcID   string         `json:"proc_id,omitempty"`   // 流程ID
	Starter  string         `json:"starter,omitempty"`   // 启动人人组
	BusiKey  string         `json:"busi_key,omitempty"`  // 业务单据ID
	BusiType string         `json:"busi_type,omitempty"` // 业务单据类型
	Input    map[string]any `json:"input,omitempty"`     // 启动参数 map[string]any
}

type UserTaskCommitCmd struct {
	TaskID   string         `json:"task_id,omitempty"`  // 任务ID
	Operator string         `json:"operator,omitempty"` // 任务提交人，对应的人组
	Input    map[string]any `json:"input,omitempty"`    // 提交参数，map[string]any
	Result   int            `json:"result,omitempty"`   // 任务执行的结果
	Version  int            `json:"version,omitempty"`
}

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

	e.isRunning = true
	return
}

func (e *Engine) StartProc(ctx context.Context, cmd StartProcCmd) (string, error) {
	if e.Templates.GetTemplate(cmd.ProcID) == nil {
		return "", fmt.Errorf("未找到流程(ProcID = %s)", cmd.ProcID)
	}

	instID := e.NewID()

	inst := &ProcInst{Engine: e, InstID: instID}

	start := &StartEventOp{
		Exec: Exec{
			ProcInst: inst,
		},
		cmd: cmd,
	}
	return instID, e.Emit(start)
}

func (e *Engine) CommitTask(ctx context.Context, cmd UserTaskCommitCmd) error {
	ut := UserTask{}
	if err := e.LoadUserTask(ctx, cmd.TaskID, &ut); err != nil {
		return fmt.Errorf("未找任务:%s > %w", cmd.TaskID, err)
	}

	if ut.Exec.ExecID != "" {
		if err := e.LoadExec(ctx, ut.Exec.ExecID, &ut.Exec); err != nil {
			return fmt.Errorf("未找执行:%s > %w", ut.Exec.ExecID, err)
		}
	}

	inst := &ProcInst{Engine: e}
	if err := e.LoadProcInst(ctx, ut.InstID, inst); err != nil {
		return fmt.Errorf("未找到流程实例:%s > %w", ut.InstID, err)
	}
	ut.ProcInst = inst
	u := &userTaskRunOp{UserTask: ut, cmd: cmd}
	return e.Emit(u)
}

func (e *Engine) Emit(ops ...Activity) error {
	if !e.isRunning {
		return errors.New("引擎未运行")
	}
	return e.driver.Emit(ops...)
}
