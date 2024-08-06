package loong

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/it512/loong/bpmn"
	"github.com/it512/loong/bpmn/zeebe"
)

const begin_version = 7

type UserTask struct {
	Variable `json:"-"`

	TaskID  string `json:"task_id,omitempty"`
	FormKey string `json:"form_key,omitempty"`

	ActName string `json:"act_name,omitempty"`
	ActID   string `json:"act_id,omitempty"`

	Assignee        string   `json:"assignee,omitempty"`
	CandidateGroups []string `json:"candidate_groups,omitempty"`
	CandidateUsers  []string `json:"candidate_users,omitempty"`

	Owner    string `json:"owner,omitempty"`
	Operator string `json:"operator,omitempty"`

	Result int `json:"result,omitempty"`

	BatchNo string `json:"batch_no,omitempty"`

	Status int `json:"status,omitempty"`

	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`

	Version int
}

type userTaskOp struct {
	InOut
	UserTask
	bpmn.TUserTask
}

func (u userTaskOp) Type() ActivityType                               { return AT_USER_TASK }
func (u userTaskOp) GetTaskDefinition(context.Context) TaskDefinition { return u.Type() }
func (userTaskOp) Emit(ctx context.Context, emt Emitter) error        { return nil }

func (c *userTaskOp) Do(ctx context.Context) error {
	if err := io(ctx, c, c); err != nil {
		return err
	}

	c.UserTask.FormKey = c.TUserTask.FormDefinition.FormKey
	c.UserTask.InstID = c.ProcInst.InstID
	c.UserTask.ActID = c.TUserTask.GetId()
	c.UserTask.ActName = c.TUserTask.GetName()
	c.UserTask.StartTime = time.Now()
	c.UserTask.Status = STATUS_START
	c.UserTask.BatchNo = c.Engine.NewID()
	c.UserTask.Version = begin_version
	c.UserTask.Owner = c.ProcInst.Starter

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

			a.Variable.PutInput("loopCounter", i)
			if key := milc.GetInputElement(); key != "" {
				a.Variable.PutInput(key, item)
			}
			if a.Assignee, a.CandidateGroups, a.CandidateUsers, err = assign(ctx, a.Variable, ad); err != nil {
				return err
			}
			tasks = append(tasks, a)
		}
	} else {
		var err error
		var a UserTask = c.UserTask // copy
		a.TaskID = c.Engine.NewID()

		if a.Assignee, a.CandidateGroups, a.CandidateUsers, err = assign(ctx, a.Variable, ad); err != nil {
			return err
		}
		tasks = append(tasks, a)
	}

	if err := c.Variable.saveTo(ctx, c.Storer); err != nil {
		return err
	}
	return c.Storer.CreateTasks(ctx, tasks...)
}

func groups(ctx context.Context, ae ActivationEvaluator, str string) (result []string, err error) {
	if str == "" {
		return
	}

	var (
		s  string
		ok bool
		a  any
	)
	if s, ok = exp(str); ok {
		if a, err = ae.Eval(ctx, str); err != nil {
			return
		}

		switch v := a.(type) {
		case []string:
			result = v
		case string:
			if v != "" {
				result = []string{v}
			}
		default:
			panic("执行人必须为string或者[]string")
		}
		return
	}

	return zeebe.SplitTrim(s, ","), nil
}

func assignee(ctx context.Context, ae ActivationEvaluator, str string) (result string, err error) {
	if result = str; str == "" {
		return
	}

	if _, ok := exp(str); ok {
		result, _, err = eval[string](ctx, ae, str)
	}

	return
}

func assign(ctx context.Context, ae ActivationEvaluator, ad zeebe.TAssignmentDefinition) (a string, b []string, c []string, err error) {
	if a, err = assignee(ctx, ae, ad.Assignee); err != nil {
		return
	}
	if b, err = groups(ctx, ae, ad.CandidateGroups); err != nil {
		return
	}
	if c, err = groups(ctx, ae, ad.CandidateUsers); err != nil {
		return
	}

	return
}

type UserTaskCommitCmd struct {
	InstID   string         `json:"inst_id,omitempty"`  // 实例ID
	TaskID   string         `json:"task_id,omitempty"`  // 任务ID
	Operator string         `json:"operator,omitempty"` // 任务提交人，对应的人组
	Input    map[string]any `json:"input,omitempty"`    // 提交参数，map[string]any
	Var      map[string]any `json:"var,omitempty"`      // 提交参数，map[string]any
	Result   int            `json:"result,omitempty"`   // 任务执行的结果
	Version  int            `json:"version,omitempty"`

	InOut
	UserTask
	bpmn.TUserTask
}

func (u UserTaskCommitCmd) Type() ActivityType                               { return AT_USER_TASK_COMMIT }
func (u UserTaskCommitCmd) GetTaskDefinition(context.Context) TaskDefinition { return u.Type() }

func (c UserTaskCommitCmd) check() error {
	if c.InstID == "" {
		return errors.New("参数InstID为空")
	}

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

	c.Exec.ProcInst = &ProcInst{Engine: e, Var: NewVar()}
	if err := e.LoadProcInst(ctx, c.InstID, c.Exec.ProcInst); err != nil {
		return fmt.Errorf("未找到流程实例:%s > %w", c.UserTask.InstID, err)
	}

	if c.Exec.ProcInst.Status != STATUS_START {
		return fmt.Errorf("流程实例: %s 当前状态为: %d, 已经终止", c.UserTask.InstID, c.ProcInst.Status)
	}

	c.UserTask.TaskID = c.TaskID
	c.UserTask.Version = c.Version
	if err := e.LoadUserTask(ctx, c.TaskID, &c.UserTask); err != nil {
		return fmt.Errorf("未找任务:%s > %w", c.TaskID, err)
	}

	if c.Exec.ExecID != "" {
		if err := e.LoadExec(ctx, c.UserTask.InstID, c.Exec.ExecID, &c.Exec); err != nil {
			return fmt.Errorf("未找执行:%s > %w", c.Exec.ExecID, err)
		}
	}

	c.Exec.ProcInst.Template = e.GetTemplate(c.ProcID)

	var ok bool
	if c.TUserTask, ok = c.Template.FindUserTask(c.UserTask.ActID); !ok {
		return fmt.Errorf("未找到环节:%s", c.UserTask.ActID)
	}

	if len(c.Var) > 0 {
		c.Variable.ProcInst.Var = Merge(c.Variable.ProcInst.Var, c.Var)
		c.Variable.isVarChanged = true
	}

	c.InOut = newInOut()

	c.Variable.Param = NewVar().
		Put("operator", c.Operator).
		Put("result", c.Result).
		Put("act_id", c.TUserTask.Id).
		Put("task_id", c.TaskID).
		Put("form_key", c.FormKey)

	c.Variable.Input = Merge(c.Variable.Input, c.Input)

	return nil
}

func (c *UserTaskCommitCmd) Do(ctx context.Context) error {
	if err := io(ctx, c, c); err == nil {
		if err = c.Variable.saveTo(ctx, c.Storer); err != nil {
			return err
		}
	}

	c.UserTask.Status = STATUS_FINISH
	c.UserTask.EndTime = time.Now()
	c.UserTask.Operator = c.Operator
	c.UserTask.Result = c.Result
	return c.Storer.EndUserTask(ctx, c.UserTask)
}

func (t *UserTaskCommitCmd) Emit(ctx context.Context, emt Emitter) error {
	var v Variable
	v.Exec = t.Variable.Exec
	v.Input = Merge(v.Input, t.Input) // 丢弃掉原来的Input，并将新的Input传递下去(bug fix #5)

	if !t.TUserTask.HasMultiInstanceLoopCharacteristics() {
		return emt.Emit(fromOuter(ctx, v, t))
	}

	milc := t.TUserTask.GetMultiInstanceLoopCharacteristics()

	ut, err := t.Storer.LoadUserTaskBatch(ctx, t.UserTask.BatchNo)
	if err != nil {
		return err
	}

	vote := newVote(ut, t.Engine, t.Variable.Input)

	var pass bool
	if pass, err = vote.Test(ctx, milc.GetCompletionCondition()); err != nil {
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
		if a, err = vote.Eval(ctx, milc.GetOutputElement()); err != nil {
			return err
		}
		v.PutInput(milc.GetOutputCollection(), a)
	}

	return emt.Emit(fromOuter(ctx, v, t))
}
