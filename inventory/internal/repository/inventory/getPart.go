package inventory

import (
	"context"
	"errors"
	"fmt"

	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	var part repoModel.Part
	err := r.collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&part)
	fmt.Println(part)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Part{}, model.ErrPartNotFound
		}
		return model.Part{}, err
	}

	return converter.PartToModel(part), nil
}
