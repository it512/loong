package mongox

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	client *mongo.Client

	db    *mongo.Database
	instC *mongo.Collection
	taskC *mongo.Collection
	execC *mongo.Collection
}
