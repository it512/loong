package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

// fork
func (m *Store) ForkExec(ctx context.Context, xs []loong.Exec) error {
	/*
		a := InterfaceSlice(xs)
		_, err := m.ExecColl().InsertMany(ctx, a)
		return err
	*/
	return nil
}

// join
func (m *Store) JoinExec(ctx context.Context, ex *loong.Exec) error {
	f := bson.D{{Key: "exec_id", Value: ex.ExecID}}
	update := bson.D{{Key: "$set", Value: bson.M{"status": ex.Status, "join_tag": ex.JoinTag}}}
	_, err := m.ExecColl().UpdateOne(ctx, f, update)
	return err
}

func (m *Store) LoadForks(ctx context.Context, instID, forkID string) ([]loong.Exec, error) {
	return nil, nil
}

func (m *Store) LoadExec(ctx context.Context, instID, execID string, ex *loong.Exec) error {
	f := bson.D{{Key: "exec_id", Value: ex.ExecID}, {Key: "inst_id", Value: instID}}
	sr := m.ExecColl().FindOne(ctx, f)
	return sr.Decode(ex)
}
