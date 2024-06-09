package mongox

import (
	"context"
	"time"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Store) CreateTasks(ctx context.Context, tasks ...loong.UserTask) error {
	a := make([]any, len(tasks))
	_ = loong.Each(tasks, func(ut loong.UserTask, i int) error {
		a[i] = userTaskConv(ut)
		return nil
	})
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
				Value: bson.M{"status": loong.STATUS_FINISH, "end_time": time.Now(), "operator": loong.Robot01},
			},
		},
	)
	return err
}
