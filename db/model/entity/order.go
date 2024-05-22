package entity

type OrderEntity struct {
	OrderID     int    `json:"order_id"`
	UserIDFK    int    `json:"user_id_fk"`
	ProductIDFK int    `json:"product_id_fk"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
	CreatedAt   string `json:"created_at"`
}
