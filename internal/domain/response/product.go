package response

import "goods-square/internal/domain/model"

type Products struct {
	Products model.Products `json:"products"`
}
