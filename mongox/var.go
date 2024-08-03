package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Store) SaveVar(ctx context.Context, procInst *loong.ProcInst) error {
	_, err := m.InstColl().UpdateOne(ctx,
		bson.D{
			{Key: "inst_id", Value: procInst.InstID},
		},

		bson.D{
			{Key: "$set",
				Value: bson.M{
					"var": procInst.Var,
				},
			},
		},
	)
	return err
}
