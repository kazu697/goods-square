package repository

import (
	"context"
	"goods-square/internal/domain/model"
	"goods-square/internal/domain/repository"
	"goods-square/internal/infrastracture/dto"

	"github.com/jinzhu/gorm"
)

type ProductRepo struct{}

func (r *ProductRepo) Get(ctx context.Context, db *gorm.DB, query *repository.ProductGetQuery) (*model.Product, error) {

	p := dto.Product{ID: query.ProductID}
	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p.ToEntity(), nil
}

func NewProduct() repository.Product {
	return &ProductRepo{}
}
