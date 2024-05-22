package web

type OrderRequest struct {
	Quantity   int `validate:"required" json:"quantity"`
	TotalPrice int `validate:"required" json:"total_price"`
}
