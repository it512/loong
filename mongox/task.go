package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func x(ut loong.UserTask) bson.M {
	return bson.M{
		"task_id": ut.TaskID,
		"inst_id": ut.InstID,
		"exec_id": ut.ExecID,

		"busi_key":  ut.BusiKey,
		"busi_type": ut.BusiType,

		"form_key": ut.FormKey,
		"act_id":   ut.ActID,
		"act_name": ut.ActName,

		"assignee":         ut.Assignee,
		"candidate_users":  ut.CandidateUsers,
		"candidate_groups": ut.CandidateGroups,

		"input": ut.Exec.Input,

		"batch_no": ut.BatchNo,

		"start_time": ut.StartTime,
		"status":     ut.Status,

		"version": ut.Version,
	}
}

func (m *Store) CreateTasks(ctx context.Context, tasks ...loong.UserTask) error {
	a := make([]any, len(tasks))
	for i, t := range tasks {
		a[i] = x(t)
	}
	_, err := m.TaskColl().InsertMany(ctx, a)
	return err
}

func (m *Store) EndUserTask(ctx context.Context, ut loong.UserTask) error {
	_, err := m.TaskColl().UpdateOne(ctx, bson.D{{Key: "task_id", Value: ut.TaskID}},
		bson.D{
			{Key: "$set",
				Value: bson.M{
					"status":   ut.Status,
					"end_time": ut.EndTime,
					"operator": ut.Operator,
					"input2":   ut.Exec.Input,
				},
			},
		},
	)
	return err
}

func (m *Store) LoadUserTask(ctx context.Context, taskID string, ut *loong.UserTask) error {
	sr := m.TaskColl().FindOne(ctx, bson.D{{Key: "task_id", Value: taskID}})
	return sr.Decode(ut)
}

func (m *Store) LoadUserTaskBatch(ctx context.Context, batchNO string) ([]loong.UserTask, error) {
	c, err := m.TaskColl().Find(ctx, bson.D{{Key: "batch_no", Value: batchNO}})
	if err != nil {
		return nil, err
	}

	var uts []loong.UserTask
	err = c.All(ctx, &uts)

	return uts, err
}

func (m *Store) EndUserTaskBatch(ctx context.Context, batchNO string) error {
	_, err := m.TaskColl().UpdateMany(ctx,
		bson.D{
			{Key: "batch_no", Value: batchNO},
			{Key: "status", Value: loong.STATUS_START},
		},

		bson.D{
			{Key: "$set",
				Value: bson.M{"status": loong.STATUS_FINISH},
			},
		},
	)
	return err
}
