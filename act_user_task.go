package loong

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"time"

	"github.com/it512/loong/bpmn"
	"github.com/it512/loong/bpmn/zeebe"
)

const begin_version = 7

type UserTask struct {
	Exec `json:"-"`

	TaskID  string `json:"task_id,omitempty"`
	FormKey string `json:"form_key,omitempty"`

	ActName string `json:"act_name,omitempty"`
	ActID   string `json:"act_id,omitempty"`

	Assignee        string `json:"assignee,omitempty"`
	CandidateGroups string `json:"candidate_groups,omitempty"`
	CandidateUsers  string `json:"candidate_users,omitempty"`

	Operator string `json:"operator,omitempty"`

	Result int `json:"result,omitempty"`

	BatchNo string `json:"batch_no,omitempty"`

	Status int `json:"status,omitempty"`

	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`

	Version int
}

type userTaskOp struct {
	UserTask
	bpmn.TUserTask
	InOut

	UnimplementedActivity
}

func (c *userTaskOp) Do(ctx context.Context) error {
	/*
		if err := io(ctx, c, c.Exec.Input); err != nil {
			return err
		}
	*/

	c.UserTask.FormKey = c.TUserTask.FormDefinition.FormKey
	c.UserTask.InstID = c.ProcInst.InstID
	c.UserTask.ActID = c.TUserTask.GetId()
	c.UserTask.ActName = c.TUserTask.GetName()
	c.UserTask.StartTime = time.Now()
	c.UserTask.Status = STATUS_START
	c.UserTask.BatchNo = c.Engine.NewID()
	c.UserTask.Version = begin_version

	ad := c.TUserTask.AssignmentDefinition

	var tasks []UserTask
	if c.TUserTask.HasMultiInstanceLoopCharacteristics() {
		milc := c.TUserTask.GetMultiInstanceLoopCharacteristics()
		items, _, err := eval[[]string](ctx, c, milc.GetInputCollection())
		if err != nil {
			return err
		}

		for i, item := range items {
			var a UserTask = c.UserTask // copy
			a.TaskID = c.Engine.NewID()

			a.Exec.Input = maps.Clone(c.Exec.Input)
			a.Exec.Input.Set("loopCounter", i)
			if key := milc.GetInputElement(); key != "" {
				a.Exec.Input.Set(key, item)
			}
			if a.Assignee, a.CandidateGroups, a.CandidateUsers, err = assign(ctx, a.Exec, ad); err != nil {
				return err
			}
			tasks = append(tasks, a)
		}
	} else {
		var err error
		var a UserTask = c.UserTask // copy
		a.TaskID = c.Engine.NewID()

		if a.Assignee, a.CandidateGroups, a.CandidateUsers, err = assign(ctx, a.Exec, ad); err != nil {
			return err
		}
		tasks = append(tasks, a)
	}
	return c.Storer.CreateTasks(ctx, tasks...)
}

func assign(ctx context.Context, ae ActivationEvaluator, ad zeebe.TAssignmentDefinition) (a string, b string, c string, err error) {
	if a, _, err = eval[string](ctx, ae, ad.Assignee); err != nil {
		panic(fmt.Errorf("执行人为空>%w", err))
	}
	if b, _, err = eval[string](ctx, ae, ad.CandidateGroups); err != nil {
		panic(fmt.Errorf("执行人为空>%w", err))
	}
	if c, _, err = eval[string](ctx, ae, ad.CandidateUsers); err != nil {
		panic(fmt.Errorf("执行人为空>%w", err))
	}
	return
}

type UserTaskCommitCmd struct {
	TaskID   string         `json:"task_id,omitempty"`  // 任务ID
	Operator string         `json:"operator,omitempty"` // 任务提交人，对应的人组
	Input    map[string]any `json:"input,omitempty"`    // 提交参数，map[string]any
	Result   int            `json:"result,omitempty"`   // 任务执行的结果
	Version  int            `json:"version,omitempty"`

	UserTask
	bpmn.TUserTask

	UnimplementedActivity
}

func (c UserTaskCommitCmd) check() error {
	if c.TaskID == "" {
		return errors.New("参数TaskID为空")
	}

	if c.Operator == "" {
		return errors.New("参数Operator为空")
	}

	if c.Version < begin_version {
		return errors.New("version错误")
	}

	return nil
}

func (c *UserTaskCommitCmd) Bind(ctx context.Context, e *Engine) error {
	if err := c.check(); err != nil {
		return err
	}

	c.Exec.ProcInst = &ProcInst{Engine: e}

	c.UserTask.TaskID = c.TaskID
	c.UserTask.Version = c.Version
	if err := e.LoadUserTask(ctx, c.TaskID, &c.UserTask); err != nil {
		return fmt.Errorf("未找任务:%s > %w", c.TaskID, err)
	}

	if c.Exec.ExecID != "" {
		if err := e.LoadExec(ctx, c.Exec.ExecID, &c.Exec); err != nil {
			return fmt.Errorf("未找执行:%s > %w", c.Exec.ExecID, err)
		}
	}

	if err := e.LoadProcInst(ctx, c.UserTask.InstID, c.Exec.ProcInst); err != nil {
		return fmt.Errorf("未找到流程实例:%s > %w", c.UserTask.InstID, err)
	}

	c.Exec.ProcInst.Template = e.GetTemplate(c.ProcID)

	var ok bool
	if c.TUserTask, ok = c.Template.FindUserTask(c.UserTask.ActID); !ok {
		return fmt.Errorf("未找到环节:%s", c.UserTask.ActID)
	}

	c.Exec.Input = Merge(c.Exec.Input, c.Input)
	return nil
}

func (c *UserTaskCommitCmd) Do(ctx context.Context) error {
	c.UserTask.Status = STATUS_FINISH
	c.UserTask.EndTime = time.Now()
	c.UserTask.Operator = c.Operator
	c.UserTask.Result = c.Result
	return c.Storer.EndUserTask(ctx, c.UserTask)
}

func (t *UserTaskCommitCmd) Emit(ctx context.Context, emt Emitter) error {
	if !t.TUserTask.HasMultiInstanceLoopCharacteristics() {
		return emt.Emit(fromOuter(ctx, t.Exec, t))
	}

	milc := t.TUserTask.GetMultiInstanceLoopCharacteristics()

	ut, err := t.Storer.LoadUserTaskBatch(ctx, t.UserTask.BatchNo)
	if err != nil {
		return err
	}

	v := newVote(ut, t.Engine, t.Exec.Input)

	var pass bool
	if pass, err = v.Test(ctx, milc.GetCompletionCondition()); err != nil {
		return err
	}

	if !pass { // 投票不通过
		return nil
	}

	if err = t.Storer.EndUserTaskBatch(ctx, t.UserTask.BatchNo); err != nil {
		return err
	}

	if milc.GetOutputCollection() != "" {
		var a any
		if a, err = v.Eval(ctx, milc.GetOutputElement()); err != nil {
			return err
		}
		t.Exec.Input.Set(milc.GetOutputCollection(), a)
	}

	return emt.Emit(fromOuter(ctx, t.Exec, t))
}
