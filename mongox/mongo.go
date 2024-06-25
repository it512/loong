package mongox

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (s *Store) DoTx(ctx context.Context, fn func(context.Context) error) error {
	sess, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)
	_, err = sess.WithTransaction(ctx, func(sc mongo.SessionContext) (any, error) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				if err = sc.AbortTransaction(ctx); err != nil {
					log.Println(err)
				}
			}
		}()
		return nil, fn(sc)
	})
	return err
}

/*
func (s *Store) DoTrans2(ctx context.Context, fn func(loong.TxContext) error) error {
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
*/
