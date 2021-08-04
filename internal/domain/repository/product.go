package repository

import (
	"context"

	"goods-square/internal/domain/model"

	"github.com/jinzhu/gorm"
)

type Product interface {
	Get(ctx context.Context, db *gorm.DB, query *ProductGetQuery) (*model.Product, error)
}

type ProductGetQuery struct {
	ProductID int
}
