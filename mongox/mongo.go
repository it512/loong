package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDB(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).
		SetBSONOptions(
			&options.BSONOptions{
				UseJSONStructTags: true,
			},
		))
	return client, err
}

func MongoStore(client *mongo.Client) loong.Option {
	store := NewStore(client)
	return loong.SetStore(store)
}

type Store struct {
	client *mongo.Client
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		client: client,
	}
}

func (s *Store) InstColl() *mongo.Collection {
	return s.client.Database("loong").Collection("inst")
}

func (s *Store) ExecColl() *mongo.Collection {
	return s.client.Database("loong").Collection("exec")
}

func (s *Store) TaskColl() *mongo.Collection {
	return s.client.Database("loong").Collection("task")
}

func InterfaceSlice[T any](slice []T) []any {
	res := make([]any, len(slice))
	for i, v := range slice {
		res[i] = v
	}
	return res
}
