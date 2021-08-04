package handler

import (
	"goods-square/internal/domain/model"
	"goods-square/internal/service"
	"goods-square/pkg/renderer"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	svc service.Product
}

func (h *ProductHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		return
	}

	res, err := h.svc.Get(ctx, &service.ProductGetRequest{
		ProductID: id,
	})

	renderer.RenderJson(w, http.StatusOK, &model.Product{
		ID:   res.ID,
		Name: res.Name,
		Type: res.Type,
	})
}

func NewProductHandler(
	svc service.Product,
) *ProductHandler {
	return &ProductHandler{
		svc,
	}
}
