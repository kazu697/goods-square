package service

import (
	"context"
	"goods-square/internal/domain/model"
	"goods-square/internal/domain/repository"

	"github.com/jinzhu/gorm"
)

type ProductService struct {
	db         *gorm.DB
	repository repository.Product
}

func (s *ProductService) Get(
	ctx context.Context,
	param *ProductGetRequest,
) (*model.Product, error) {
	q := repository.ProductGetQuery{
		ProductID: param.ProductID,
	}

	res, err := s.repository.Get(ctx, s.db, &q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProductService(
	db *gorm.DB,
	repository repository.Product,
) Product {
	return &ProductService{
		db,
		repository,
	}
}
