package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Store) LoadProcInst(ctx context.Context, instID string, p *loong.ProcInst) error {
	sr := m.InstColl().FindOne(ctx, bson.D{{Key: "inst_id", Value: instID}})
	return sr.Decode(p)
}

func (m *Store) CreateProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	_, err := m.InstColl().InsertOne(ctx, procInst)
	return err
}

func (m *Store) EndProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	_, err := m.InstColl().UpdateOne(ctx, bson.D{{Key: "inst_id", Value: procInst.InstID}},
		bson.D{
			{Key: "$set",
				Value: bson.M{
					"status":   procInst.Status,
					"end_time": procInst.EndTime,
				},
			},
		},
	)
	return err
}
