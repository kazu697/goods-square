package repository

import (
	"context"
	"goods-square/internal/domain/model"
	"goods-square/internal/domain/repository"
	"goods-square/internal/infrastracture/dto"

	"github.com/jinzhu/gorm"
)

type ProductRepo struct{}

func (r *ProductRepo) List(ctx context.Context, db *gorm.DB, query *repository.ProductListQuery) (model.Products, error) {
	if query.Name != nil {
		db = db.Where("products.name = ?", query.Name)
	}
	if query.Type != nil {
		db = db.Where("products.type = ?", query.Type)
	}

	ps := dto.Products{}

	err := db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps.ToEntities(), nil

}

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
