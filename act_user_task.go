package loong

import (
	"context"
	"fmt"
	"maps"
	"time"

	"github.com/it512/loong/bpmn"
	"github.com/it512/loong/bpmn/zeebe"
)

type UserTask struct {
	Exec

	TaskID  string
	InstID  string
	FormKey string

	ActName string
	ActID   string

	Assignee        string
	CandidateGroups string
	CandidateUsers  string

	Operator string

	Result int

	BatchNo string

	Status int

	StartTime time.Time
	EndTime   time.Time
}

type userTaskOp struct {
	UserTask
	bpmn.TUserTask
	InOut
}

func (c *userTaskOp) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return idTaskDef{c.GetId()}
}

func (c *userTaskOp) Do(ctx context.Context) error {
	io(ctx, c, c.Exec.Input)

	c.UserTask.FormKey = c.TUserTask.FormDefinition.FormKey
	c.UserTask.InstID = c.ProcInst.InstID
	c.UserTask.ActID = c.TUserTask.GetId()
	c.UserTask.ActName = c.TUserTask.GetName()
	c.UserTask.StartTime = time.Now()
	c.UserTask.Status = STATUS_START
	c.UserTask.BatchNo = c.IDGen.NewID()

	ad := c.TUserTask.AssignmentDefinition

	var tasks []UserTask
	if c.TUserTask.HasMultiInstanceLoopCharacteristics() {
		milc := c.TUserTask.GetMultiInstanceLoopCharacteristics()
		items, _, err := eval[[]string](ctx, c, milc.GetInputCollection())
		if err != nil {
			return err
		}

		for _, item := range items {
			var a UserTask = c.UserTask // copy
			a.Exec.Input = maps.Clone(c.Exec.Input)
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
		if a.Assignee, a.CandidateGroups, a.CandidateUsers, err = assign(ctx, a.Exec, ad); err != nil {
			return err
		}
		tasks = append(tasks, a)
	}
	return c.Store.CreateTasks(ctx, tasks...)
}

func assign(ctx context.Context, ae ActivationEvaluator, ad zeebe.TAssignmentDefinition) (a string, b string, c string, err error) {
	if a, _, err = eval[string](ctx, ae, ad.Assignee); err != nil {
		panic(fmt.Errorf("执行人为空:%w", err))
	}
	if b, _, err = eval[string](ctx, ae, ad.CandidateGroups); err != nil {
		panic(fmt.Errorf("执行人为空:%w", err))
	}
	if c, _, err = eval[string](ctx, ae, ad.CandidateUsers); err != nil {
		panic(fmt.Errorf("执行人为空:%w", err))
	}
	return
}

type userTaskRunOp struct {
	UserTask
	bpmn.TUserTask

	cmd UserTaskCommitCmd
}

func (c *userTaskRunOp) Do(ctx context.Context) error {
	c.ProcInst.Template = c.GetTemplate(c.ProcID)

	var ok bool
	if c.TUserTask, ok = c.Template.FindUserTask(c.UserTask.ActID); !ok {
		panic("未找到环节")
	}

	c.Exec.Input = Merge(c.Exec.Input, c.cmd.Input)

	c.UserTask.Status = STATUS_FINISH
	c.UserTask.EndTime = time.Now()
	c.UserTask.Operator = c.cmd.Operator
	c.UserTask.Result = c.cmd.Result
	return c.Store.EndUserTask(ctx, c.UserTask)
}

func (t *userTaskRunOp) Emit(ctx context.Context, emt Emitter) error {
	if !t.TUserTask.HasMultiInstanceLoopCharacteristics() {
		return t.EmitDefault(ctx, t.TUserTask, emt)
	}

	milc := t.TUserTask.GetMultiInstanceLoopCharacteristics()

	ut, err := t.Store.LoadUserTaskBatch(ctx, t.UserTask.BatchNo)
	if err != nil {
		return err
	}

	v := newVote(ut, t.Engine)

	var pass bool
	if pass, err = v.Test(ctx, milc.GetCompletionCondition()); err != nil {
		return err
	}

	if !pass { // 投票不通过
		return nil
	}

	if err = t.Store.EndUserTaskBatch(ctx, t.UserTask.BatchNo); err != nil {
		return err
	}

	return t.EmitDefault(ctx, t.TUserTask, emt)
}
