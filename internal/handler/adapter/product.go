package adapter

type ProductListRequestQuery struct {
	ProductID *int    `schema:"product_id"`
	Name      *string `schema:"name"`
	Type      *string `schema:"type"`
}
