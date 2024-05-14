package loong

import (
	"context"
	"time"

	"github.com/it512/loong/bpmn"
)

type UserTask struct {
	Exec

	TaskID  string
	InstID  string
	FormKey string

	ActName string
	ActID   string

	CandidateGroups string

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
}

func (c *userTaskOp) Do(ctx context.Context) error {

	c.UserTask.FormKey = c.TUserTask.FormDefinition.FormKey
	c.UserTask.InstID = c.ProcInst.InstID
	c.UserTask.ActID = c.TUserTask.GetId()
	c.UserTask.ActName = c.TUserTask.GetName()
	c.UserTask.StartTime = time.Now()
	c.UserTask.Status = STATUS_START

	var tasks []UserTask
	if c.TUserTask.HasMultiInstanceLoopCharacteristics() {
	} else {
		groupAny, err := c.Eval(ctx, c.TUserTask.AssignmentDefinition.CandidateGroups)
		if err != nil {
			return err
		}

		var ok bool
		c.UserTask.CandidateGroups, ok = groupAny.(string)
		if !ok || c.UserTask.CandidateGroups == "" {
			panic("执行人为空")
		}

		c.UserTask.BatchNo = c.IDGen.NewID()

		tasks = append(tasks, c.UserTask)
	}

	return c.Store.CreateTasks(ctx, tasks...)
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

	v := &vote{}

	ut, _ := t.Store.LoadUserTaskBatch(ctx, t.UserTask.BatchNo)
	v.Put(ut)
	a, _ := t.Engine.Evaluator.Eval(ctx, "", v.ToVar())
	b := a.(bool)
	if b {
		return t.EmitDefault(ctx, t.TUserTask, emt)
	}

	if v.numberOfActiveInstances == 0 {
		panic("")
	}
	return nil
}

type vote struct {
	numberOfInstances           int //The number of instances created.
	numberOfActiveInstances     int // The number of instances currently active.
	numberOfCompletedInstances  int // The number of instances already completed.
	numberOfTerminatedInstances int // The number of instances already terminated.
}

func (v *vote) Put(ut []UserTask) {
	for _, u := range ut {
		v.numberOfInstances = v.numberOfInstances + 1

		if u.Status == STATUS_START { // 未投票的
			v.numberOfActiveInstances = v.numberOfActiveInstances + 1
		}

		if u.Result != 0 { // 投票不通过
			v.numberOfTerminatedInstances = v.numberOfTerminatedInstances + 1
		}
	}

	// 投票通过的 = 已投票 - 投票不通过的 = 总数 - 未投票的 - 投票不通过的
	v.numberOfCompletedInstances = v.numberOfInstances - v.numberOfActiveInstances - v.numberOfTerminatedInstances
}

func (v vote) ToVar() Var {
	return NewVar()
}
