package service

import (
	"context"
	"goods-square/internal/domain/model"
)

type Product interface {
	Get(context.Context, *ProductGetRequest) (*model.Product, error)
}

type ProductGetRequest struct {
	ProductID int
}
