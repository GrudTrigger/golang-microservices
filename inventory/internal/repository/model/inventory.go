package model

import (
	"time"
)

type Category int32

type Dimensions struct {
	Length float64 `bson:"length"`
	Width  float64 `bson:"width"`
	Height float64 `bson:"height"`
	Weight float64 `bson:"weight"`
}

type Manufacturer struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website"`
}

type Part struct {
	Uuid          string         `bson:"_id,omitempty"`
	Name          string         `bson:"name"`
	Description   string         `bson:"description"`
	Price         float32        `bson:"price"`
	StockQuantity int64          `bson:"stockQuantity"`
	Category      Category       `bson:"category"`
	Dimensions    *Dimensions    `bson:"dimensions"`
	Manufacturer  *Manufacturer  `bson:"manufacturer"`
	Tags          []string       `bson:"tags"`
	Metadata      map[string]any `bson:"metadata"`
	CreatedAt     *time.Time     `bson:"created_at"`
	UpdatedAt     *time.Time     `bson:"updated_at,omitempty"`
}
