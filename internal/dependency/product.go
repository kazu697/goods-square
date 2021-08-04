package dependency

import (
	"goods-square/internal/handler"
	"goods-square/internal/infrastracture/repository"
	"goods-square/internal/routes"
	"goods-square/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/gorm"
)

func NewProduct(r *chi.Mux, db *gorm.DB) {
	repo := repository.NewProduct()
	ProductService := service.NewProductService(db, repo)
	ProductHandler := handler.NewProductHandler(ProductService)

	c := routes.Product{
		Get: ProductHandler.Get,
	}

	c.Install(r)
}
