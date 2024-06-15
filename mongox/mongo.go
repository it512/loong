package mongox

import (
	"context"

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

type Store struct {
	client *mongo.Client
	dbName string
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		client: client,
		dbName: "loong",
	}
}

func (s *Store) SetDbName(dbname string) *Store {
	s.dbName = dbname
	return s
}

func (s *Store) InstColl() *mongo.Collection {
	return s.client.Database(s.dbName).Collection("loong_inst")
}

func (s *Store) ExecColl() *mongo.Collection {
	return s.client.Database(s.dbName).Collection("loong_exec")
}

func (s *Store) TaskColl() *mongo.Collection {
	return s.client.Database(s.dbName).Collection("loon_task")
}

func InterfaceSlice[T any](slice []T) []any {
	res := make([]any, len(slice))
	for i, v := range slice {
		res[i] = v
	}
	return res
}
