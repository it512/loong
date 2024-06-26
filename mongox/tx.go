package mongox

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type TxFunc func(context.Context, func(context.Context) error) error

func (s TxFunc) DoTx(ctx context.Context, fn func(context.Context) error) error {
	return s(ctx, fn)
}

func (s *Store) Tx() TxFunc {
	return func(ctx context.Context, fn func(context.Context) error) error {
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
}

func (s *Store) NoTx() TxFunc {
	return func(ctx context.Context, fn func(context.Context) error) error {
		/*
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
		*/
		return fn(ctx)
	}
}
