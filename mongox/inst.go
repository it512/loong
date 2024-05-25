package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Store) LoadProcInst(ctx context.Context, instID string, p *loong.ProcInst) error {
	sr := m.instC.FindOne(ctx, bson.D{{Key: "inst_id", Value: instID}})
	if sr.Err() != nil {
		return sr.Err()
	}
	return sr.Decode(p)
}

func (m *Store) CreateProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	return nil
}

func (m *Store) EndProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	_, err := m.instC.UpdateOne(ctx, bson.D{{Key: "inst_id", Value: procInst.InstID}}, nil)
	return err
}
