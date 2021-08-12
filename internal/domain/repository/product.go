package repository

import (
	"context"

	"goods-square/internal/domain/model"

	"github.com/jinzhu/gorm"
)

type Product interface {
	List(ctx context.Context, db *gorm.DB, query *ProductListQuery) (model.Products, error)
	Get(ctx context.Context, db *gorm.DB, query *ProductGetQuery) (*model.Product, error)
}

type ProductListQuery struct {
	ProductID *int
	Name      *string
	Type      *string
}
type ProductGetQuery struct {
	ProductID int
}
