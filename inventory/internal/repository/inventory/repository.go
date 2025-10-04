package inventory

import "go.mongodb.org/mongo-driver/mongo"

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *repository {
	collection := db.Collection("parts")

	return &repository{
		collection,
	}
}
