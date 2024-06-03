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

func MongoStore(url string) loong.Option {
	db := loong.Must(OpenDB(url))
	store := NewStore(db)
	return loong.SetStore(store)
}

type Store struct {
	client *mongo.Client

	db *mongo.Database
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		client: client,
		db:     client.Database("loong"),
	}
}

func (s *Store) InstColl() *mongo.Collection {
	return s.client.Database("loong").Collection("inst")
}

func (s *Store) TaskColl() *mongo.Collection {
	return s.client.Database("loong").Collection("task")
}
