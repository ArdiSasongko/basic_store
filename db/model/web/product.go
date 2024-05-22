package web

type ProductRequest struct {
	Name     string `validate:"required" json:"name"`
	Quantity int    `validate:"required" json:"quantity"`
	Price    int    `validate:"required" json:"price"`
}

type ProductUpdateRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
