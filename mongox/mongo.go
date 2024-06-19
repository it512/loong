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
	return s.client.Database(s.dbName).Collection("loong_task")
}

func (s *Store) DoTrans(ctx context.Context, fn func(loong.TxContext) error) error {
	sess, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)

	return fn(&txCtx{session: sess, Context: ctx})
}

type txCtx struct {
	session mongo.Session
	context.Context
}

func (c *txCtx) Commit(ctx context.Context) error {
	return c.session.CommitTransaction(ctx)
}
func (c *txCtx) Abort(ctx context.Context) error {
	return c.session.AbortTransaction(ctx)
}
func (c *txCtx) End(ctx context.Context) {
	c.session.EndSession(ctx)
}
