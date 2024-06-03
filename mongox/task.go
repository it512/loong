package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Store) CreateTasks(ctx context.Context, tasks ...loong.UserTask) error {
	a := make([]any, len(tasks))
	for i, t := range tasks {
		a[i] = t
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
					"input":    ut.Exec.Input,
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
	return nil, nil
}

func (m *Store) EndUserTaskBatch(ctx context.Context, batchNO string) error {
	return nil
}
