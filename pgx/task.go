package pgx

import (
	"context"

	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent"
	"github.com/it512/loong/pgx/internal/ent/usertask"
)

func (m *Store) CreateTasks(ctx context.Context, tasks ...loong.UserTask) error {
	cr := m.UserTask.MapCreateBulk(tasks, func(cr *ent.UserTaskCreate, i int) {
		task := tasks[i]
		cr.
			SetInstID(task.ProcInst.InstID).
			SetExecID((task.Exec.ExecID)).
			SetBusiKey(task.ProcInst.BusiKey).
			SetBusiType(task.ProcInst.BusiType).
			SetFormKey(task.FormKey).
			SetCandidateGroups(task.CandidateGroups).
			SetActID(task.ActID).
			SetActName(task.ActName).
			SetBatchNo(task.BatchNo).
			SetInput(task.Exec.Input).
			SetStartTime(task.StartTime).
			SetStatus(task.Status)
	})

	return cr.Exec(ctx)
}

func (m *Store) EndUserTask(ctx context.Context, ut loong.UserTask) error {
	return m.UserTask.Update().Where(usertask.TaskIDEQ(ut.TaskID)).
		SetEndTime(ut.EndTime).
		SetStatus(ut.Status).
		SetOperator(ut.Operator).
		SetInput(ut.Input).
		Exec(ctx)
}

func (m *Store) LoadUserTask(ctx context.Context, taskID string, ut *loong.UserTask) error {
	q := m.UserTask.Query()
	q.Where(usertask.TaskIDEQ(taskID), usertask.StatusEQ(loong.STATUS_START))
	u, err := q.Only(ctx)
	if err != nil {
		return err
	}
	ut.TaskID = u.TaskID
	ut.InstID = u.InstID
	ut.Exec.ExecID = u.ExecID

	ut.Exec.Input = u.Input

	ut.StartTime = u.StartTime
	ut.ActID = u.ActID
	ut.CandidateGroups = u.CandidateGroups
	ut.BatchNo = u.BatchNo
	return nil
}

func (m *Store) LoadUserTaskBatch(ctx context.Context, batchNO string) ([]loong.UserTask, error) {
	return nil, nil
}
