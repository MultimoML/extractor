package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Products []Product

type Product struct {
	Id   primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Name string             `json:"name,omitempty" validate:"required"`

	IsOnPromotion   bool     `json:"is-on-promotion,omitempty" validate:"required"`
	CategoryNames   []string `json:"category-names,omitempty" validate:"required"`
	CategoryName    string   `json:"category-name,omitempty" validate:"required"`
	AllergensFilter []string `json:"allergens-filter,omitempty" validate:"required"`

	SalesUnit    string `json:"sales-unit,omitempty" validate:"required"`
	Title        string `json:"title,omitempty" validate:"required"`
	CodeInternal uint64 `json:"code-internal,omitempty" validate:"required"`

	Price     float32            `json:"price,omitempty" validate:"required"`
	CreatedAt primitive.DateTime `json:"created-at,omitempty" validate:"required"`
	BestPrice float32            `json:"best-price,omitempty" validate:"required"`

	StockStatus         string  `json:"stock-status,omitempty" validate:"required"`
	IsNew               bool    `json:"is-new,omitempty" validate:"required"`
	ImageURL            string  `json:"image-url,omitempty" validate:"required"`
	ApproxWeightProduct bool    `json:"approx-weight-product,omitempty" validate:"required"`
	URL                 string  `json:"url,omitempty" validate:"required"`
	Brand               string  `json:"brand,omitempty" validate:"required"`
	PricePerUnit        string  `json:"price-per-unit,omitempty" validate:"required"`
	RegularPrice        float32 `json:"regular-price,omitempty" validate:"required"`
	PricePerUnitNumber  float32 `json:"price-per-unit-number,omitempty" validate:"required"`
}
