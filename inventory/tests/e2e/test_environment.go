package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
)

type Category int32

// Вставляет тестовую запчасть в mongo и возвращает ее uuid
func (env *TestEnvironment) InsertTestPartData(ctx context.Context) (string, error) {
	partUuid := gofakeit.UUID()
	now := time.Now()

	dimensions := bson.M{
		"length": gofakeit.Float64(),
		"width":  gofakeit.Float64(),
		"height": gofakeit.Float64(),
		"weight": gofakeit.Float64(),
	}

	manufacturer := bson.M{
		"name":    gofakeit.Company(),
		"country": gofakeit.Country(),
		"website": gofakeit.URL(),
	}

	part := bson.M{
		"_id":           partUuid,
		"name":          "engine",
		"description":   gofakeit.Product().Description,
		"price":         gofakeit.Float32(),
		"stockQuantity": gofakeit.Int64(),
		"category":      Category(1),
		"dimensions":    dimensions,
		"manufacturer":  manufacturer,
		"tags":          []string{"a", "b", "c"},
		"createdAt":     now,
	}

	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory"
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(inventoryCollectionName).InsertOne(ctx, part)
	if err != nil {
		return "", err
	}
	return partUuid, nil
}

func (env *TestEnvironment) InsertTestSlicePartsData(ctx context.Context) error {
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory"
	}

	var parts []interface{}
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("name-%d", i)
		part := NewPart(name)
		parts = append(parts, part)
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(inventoryCollectionName).InsertMany(ctx, parts)
	if err != nil {
		return err
	}
	return nil
}

func (env *TestEnvironment) ClearPartsCollection(ctx context.Context) error {
	databaseName := os.Getenv("MONGO_DATABASE")
	if databaseName == "" {
		databaseName = "inventory"
	}

	_, err := env.Mongo.Client().Database(databaseName).Collection(inventoryCollectionName).DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}
	return nil
}

func NewPart(name string) bson.M {
	return bson.M{
		"_id":           gofakeit.UUID(),
		"name":          name,
		"description":   gofakeit.Product().Description,
		"price":         gofakeit.Float32Range(10, 1000),
		"stockQuantity": gofakeit.Int64(),
		"category":      Category(1),
		"dimensions": bson.M{
			"width":  gofakeit.Float32Range(1, 10),
			"height": gofakeit.Float32Range(1, 10),
			"depth":  gofakeit.Float32Range(1, 10),
		},
		"manufacturer": bson.M{
			"name": gofakeit.Company(),
			"city": gofakeit.City(),
		},
		"tags":      []string{"a", "b", "c"},
		"createdAt": time.Now(),
	}
}
