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
		a[i] = usertask_2_usertaskdata(ut)
		return nil
	})
	_, err := m.TaskColl().InsertMany(ctx, a)
	return err
}

func (m *Store) EndUserTask(ctx context.Context, ut loong.UserTask) error {
	_, err := m.TaskColl().UpdateOne(ctx, bson.D{
		{Key: "task_id", Value: ut.TaskID},
		{Key: "version", Value: ut.Version},
	},
		bson.D{
			{Key: "$set",
				Value: bson.M{
					"status":   ut.Status,
					"end_time": ut.EndTime,
					"operator": ut.Operator,
					"input2":   ut.Exec.Input,
				},
			},
			{Key: "$inc", Value: bson.D{{Key: "version", Value: 1}}},
		},
	)
	return err
}

func (m *Store) LoadUserTask(ctx context.Context, taskID string, ut *loong.UserTask) error {
	u := userTaskData{}
	sr := m.TaskColl().FindOne(ctx, bson.D{
		{Key: "task_id", Value: taskID},
		{Key: "version", Value: ut.Version},
	})

	if err := sr.Decode(&u); err != nil {
		return err
	}

	usertaskdata_ptr_2_usertask_ptr(ut, &u)

	return nil
}

func (m *Store) LoadUserTaskBatch(ctx context.Context, batchNO string) ([]loong.UserTask, error) {
	c, err := m.TaskColl().Find(ctx, bson.D{{Key: "batch_no", Value: batchNO}})
	if err != nil {
		return nil, err
	}

	var uts []loong.UserTask

	var data []userTaskData
	if err = c.All(ctx, &data); err != nil {
		return uts, err
	}

	uts = make([]loong.UserTask, len(data))
	err = loong.Each(data, func(u userTaskData, i int) error {
		uts[i] = usertaskdata_2_usertask_no_exec(u)
		return nil
	})

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
				Value: bson.D{
					{Key: "status", Value: loong.STATUS_FINISH},
					{Key: "end_time", Value: time.Now()},
					{Key: "operator", Value: loong.Robot},
				},
			},
			{Key: "$inc", Value: bson.D{{Key: "version", Value: 100}}},
		},
	)
	return err
}
