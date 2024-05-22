package entity

import "basic-store/db/model/domain"

type ProductEntity struct {
	ProductID  int    `json:"product_id"`
	SellerIDFK int    `json:"seller_id_fk"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	CreatedAt  string `json:"created_at"`
}

func ToProduct(product domain.Products) ProductEntity {
	return ProductEntity{
		ProductID:  product.ProductID,
		SellerIDFK: product.SellerIDFK,
		Name:       product.Name,
		Quantity:   product.Quantity,
		Price:      product.Price,
		CreatedAt:  product.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToProductList(products []domain.Products) []ProductEntity {
	var result []ProductEntity
	for _, v := range products {
		result = append(result, ToProduct(v))
	}
	return result
}
