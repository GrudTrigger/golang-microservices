package inventory

import (
	"context"
	"errors"

	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	var part repoModel.Part
	findUuid, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return model.Part{}, err
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": findUuid}).Decode(&part)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Part{}, model.ErrPartNotFound
		}
		return model.Part{}, err
	}

	return converter.PartToModel(part), nil
}
