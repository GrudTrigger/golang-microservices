package inventory

import (
	"context"
	"log"

	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	var result []repoModel.Part
	mongoFilter := bson.M{}
	if len(filter.Uuids) > 0 {
		var findUuids []primitive.ObjectID
		for _, v := range filter.Uuids {
			id, err := primitive.ObjectIDFromHex(v)
			if err != nil {
				return nil, err
			}
			findUuids = append(findUuids, id)
		}
		mongoFilter["_id"] = bson.M{"$in": findUuids}
	}

	if len(filter.Names) > 0 {
		mongoFilter["name"] = bson.M{"$in": filter.Names} // исправлено
	}
	if len(filter.Categories) > 0 {
		mongoFilter["category"] = bson.M{"$in": filter.Categories}
	}

	if len(filter.ManufacturerCountries) > 0 {
		mongoFilter["manufacturer.country"] = bson.M{"$in": filter.ManufacturerCountries}
	}

	if len(filter.Tags) > 0 {
		mongoFilter["tags"] = bson.M{"$in": filter.Tags}
	}
	cursor, err := r.collection.Find(ctx, mongoFilter)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := cursor.Close(ctx)
		if cerr != nil {
			log.Printf("failed to close cursor: %v\n", cerr)
		}
	}()

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return converter.PartToModelSlice(result), nil
}
