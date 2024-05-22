package domain

import "time"

type Products struct {
	ProductID  int    `gorm:"column:product_id;primaryKey;autoIncrement"`
	SellerIDFK int    `gorm:"column:seller_id_fk"`
	Name       string `gorm:"column:name"`
	Quantity   int    `gorm:"column:quantity"`
	Price      int    `gorm:"column:price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
