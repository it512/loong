package mongox

import (
	"context"

	"github.com/it512/loong"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NoTxStore(dbname string, client *mongo.Client) loong.Option {
	store := New(dbname, client)
	return func(c *loong.Config) {
		c.Store = store
		c.Txer = store.NoTx()
	}
}

func TxStore(dbname string, client *mongo.Client) loong.Option {
	store := New(dbname, client)
	return func(c *loong.Config) {
		c.Store = store
		c.Txer = store.Tx()
	}
}

func ClientOptions(uri string) *options.ClientOptions {
	return options.
		Client().
		ApplyURI(uri).
		SetBSONOptions(
			&options.BSONOptions{
				UseJSONStructTags: true,
			},
		)
}

func OpenDB(uri string) (*mongo.Client, error) {
	return OpenDBCtx(context.Background(), uri)
}

func OpenDBWith(ctx context.Context, op *options.ClientOptions) (*mongo.Client, error) {
	return mongo.Connect(ctx, op)
}

func OpenDBCtx(ctx context.Context, uri string) (*mongo.Client, error) {
	return mongo.Connect(ctx, ClientOptions(uri))
}

type Store struct {
	client *mongo.Client
	dbName string

	instName string
	taskName string
	execName string
}

func New(dbname string, client *mongo.Client) *Store {
	return &Store{
		client: client,
		dbName: dbname,

		instName: "loong_inst",
		taskName: "loong_task",
		execName: "loong_exec",
	}
}

func (s *Store) getColl(dbname, collname string) *mongo.Collection {
	return s.client.Database(dbname).Collection(collname)
}

func (s *Store) InstColl() *mongo.Collection {
	return s.getColl(s.dbName, s.instName)
}

func (s *Store) ExecColl() *mongo.Collection {
	return s.getColl(s.dbName, s.execName)
}

func (s *Store) TaskColl() *mongo.Collection {
	return s.getColl(s.dbName, s.taskName)
}
