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
	v := &vote{}

	ut, err := t.Store.LoadUserTaskBatch(ctx, t.UserTask.BatchNo)
	if err != nil {
		return err
	}

	v.Put(ut)
	b, _, err := eval2[bool](ctx, t.Engine, milc.GetCompletionCondition(), v.ToEnv())
	if err != nil {
		return err
	}

	if b { // 投票通过
		if err := t.Store.EndUserTaskBatch(ctx, t.UserTask.BatchNo); err != nil {
			return err
		}
		return t.EmitDefault(ctx, t.TUserTask, emt)
	}

	if v.numberOfActiveInstances == 0 {
		panic(fmt.Errorf("投票已经结束，未能达成通过条件 %s", milc.GetCompletionCondition()))
	}
	return nil
}

type vote struct {
	numberOfInstances           int // The number of instances created.
	numberOfActiveInstances     int // The number of instances currently active.
	numberOfCompletedInstances  int // The number of instances already completed.
	numberOfTerminatedInstances int // The number of instances already terminated.
}

func (v *vote) Put(ut []UserTask) {
	for _, u := range ut {
		v.numberOfInstances++

		if u.Status == STATUS_START { // 未投票的
			v.numberOfActiveInstances++
		}

		if u.Result != 0 { // 投票不通过
			v.numberOfTerminatedInstances++
		}
	}

	// 投票通过的 = 已投票 - 投票不通过的 = 总数 - 未投票的 - 投票不通过的
	v.numberOfCompletedInstances = v.numberOfInstances - v.numberOfActiveInstances - v.numberOfTerminatedInstances
}

func (v vote) ToEnv() Var {
	return NewVar().
		Put("numberOfInstances", v.numberOfInstances).
		Put("numberOfActiveInstances", v.numberOfActiveInstances).
		Put("numberOfCompletedInstances", v.numberOfCompletedInstances).
		Put("numberOfTerminatedInstances", v.numberOfTerminatedInstances)
}
