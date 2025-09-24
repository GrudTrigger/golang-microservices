package converter

import (
	"github.com/rocket-crm/inventory/internal/model"
	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
)

func PartToRepoModal(part model.Part) repoModel.Part {
	return repoModel.Part{
		Uuid: part.Uuid,                  
		Name: part.Name,                      
		Description: part.Description,                   
		Price: part.Price,      
		StockQuantity: part.StockQuantity,                
		Category: repoModel.Category(part.Category),                  
		Dimensions: (*repoModel.Dimensions)(part.Dimensions),               
		Manufacturer: (*repoModel.Manufacturer)(part.Manufacturer),           
		Tags: part.Tags,                        
		Metadata: part.Metadata,          
		CreatedAt: part.CreatedAt, 
		UpdatedAt: part.UpdatedAt,
	}
}

func PartToModel(part repoModel.Part) model.Part {
	return model.Part{
		Uuid: part.Uuid,                  
		Name: part.Name,                      
		Description: part.Description,                   
		Price: part.Price,      
		StockQuantity: part.StockQuantity,                
		Category: model.Category(part.Category),                  
		Dimensions: (*model.Dimensions)(part.Dimensions),               
		Manufacturer: (*model.Manufacturer)(part.Manufacturer),           
		Tags: part.Tags,                        
		Metadata: part.Metadata,          
		CreatedAt: part.CreatedAt, 
		UpdatedAt: part.UpdatedAt,
	}
}