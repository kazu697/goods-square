package handler

import (
	"goods-square/internal/domain/model"
	"goods-square/internal/domain/response"
	"goods-square/internal/service"
	"goods-square/pkg/renderer"
	"log"
	"net/http"
	"strconv"

	"goods-square/internal/handler/adapter"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	svc service.Product
}

func (h *ProductHandler) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	q := r.URL.Query()

	requestQuery := adapter.ProductListRequestQuery{}
	if err := queryDecoder.Decode(&requestQuery, q); err != nil {
		log.Printf(err.Error())
	}
	res, err := h.svc.List(ctx, &service.ProductListQuery{
		ProductID: requestQuery.ProductID,
		Name:      requestQuery.Name,
		Type:      requestQuery.Type,
	})
	if err != nil {
		log.Printf(err.Error())
	}

	renderer.RenderJson(w, http.StatusOK, response.Products{
		Products: res,
	})
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
