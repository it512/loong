package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) CreateInstIndex(ctx context.Context) {
	var indexes = []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "inst_id", Value: -1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "status", Value: -1}},
		},
		{
			Keys:    bson.D{{Key: "end_time", Value: -1}},
			Options: options.Index().SetExpireAfterSeconds(5 * 86400),
		},
	}

	_, _ = s.InstColl().Indexes().CreateMany(ctx, indexes)
}

func (s *Store) CreateIndex(ctx context.Context) {
	s.CreateInstIndex(ctx)
}
