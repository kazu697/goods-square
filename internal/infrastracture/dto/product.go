package dto

import "goods-square/internal/domain/model"

type Product struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"not null"`
	Type string `gorm:"not null"`
}

func NewProductFromEntity(e *model.Product) *Product {
	if e == nil {
		return nil
	}
	return &Product{
		ID:   e.ID,
		Name: e.Name,
		Type: e.Type,
	}
}

func (g *Product) ToEntity() *model.Product {
	if g == nil {
		return nil
	}
	return &model.Product{
		ID:   g.ID,
		Name: g.Name,
		Type: g.Type,
	}
}

type Products []*Product

func NewProductsFromEntities(es model.Products) Products {
	Product := make(Products, len(es))
	for idx, e := range es {
		Product[idx] = NewProductFromEntity(e)
	}
	return Product
}

func (gs Products) ToEntities() model.Products {
	Product := make(model.Products, len(gs))
	for idx, g := range gs {
		Product[idx] = g.ToEntity()
	}
	return Product
}
