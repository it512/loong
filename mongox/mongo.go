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

func (s *Store) DoTrans(ctx context.Context, fn func(loong.TxContext) error) error {
	sess, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)
	sc := mongo.NewSessionContext(ctx, sess)

	return fn(&txCtx{SessionContext: sc})
}

type txCtx struct {
	mongo.SessionContext
}

func (c *txCtx) Commit(ctx context.Context) error {
	return c.SessionContext.CommitTransaction(ctx)
}
func (c *txCtx) Abort(ctx context.Context) error {
	return c.SessionContext.AbortTransaction(ctx)
}
func (c *txCtx) End(ctx context.Context) {
	c.SessionContext.EndSession(ctx)
}
