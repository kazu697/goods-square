package service

import (
	"context"
	"goods-square/internal/domain/model"
)

type Product interface {
	List(context.Context, *ProductListQuery) (model.Products, error)
	Get(context.Context, *ProductGetRequest) (*model.Product, error)
}

type ProductListQuery struct {
	ProductID *int
	Name      *string
	Type      *string
}
type ProductGetRequest struct {
	ProductID int
}
