package domain

import "time"

type Orders struct {
	OrderID     int    `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserIDFK    int    `gorm:"column:user_id_fk"`
	ProductIDFK int    `gorm:"column:product_id_fk"`
	Name        string `gorm:"column:name"`
	Quantity    int    `gorm:"column:quantity"`
	TotalPrice  int    `gorm:"column:total_price"`
	CreatedAt   time.Time
}
