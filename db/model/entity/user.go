package entity

import "basic-store/db/model/domain"

type BuyerEntity struct {
	UserID int         `json:"user_id"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Orders interface{} `json:"orders"`
}

type SellerEntity struct {
	UserID   int         `json:"user_id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Products interface{} `json:"products"`
}

func ToBuyerHistory(user domain.Users) BuyerEntity {
	var orders []OrderEntity
	if len(user.Orders) != 0 {
		for _, v := range user.Orders {
			orders = append(orders, OrderEntity{
				OrderID:     v.OrderID,
				UserIDFK:    v.UserIDFK,
				ProductIDFK: v.ProductIDFK,
				Quantity:    v.Quantity,
				TotalPrice:  v.TotalPrice,
				CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		return BuyerEntity{
			UserID: user.UserID,
			Name:   user.Name,
			Email:  user.Email,
			Orders: orders,
		}
	}

	return BuyerEntity{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Orders: "never made a purchase",
	}
}

func ToSellerProduct(user domain.Users) SellerEntity {
	var products []ProductEntity
	if len(user.Products) != 0 {
		for _, v := range user.Products {
			products = append(products, ProductEntity{
				ProductID:  v.ProductID,
				SellerIDFK: v.SellerIDFK,
				Name:       v.Name,
				Quantity:   v.Quantity,
				Price:      v.Price,
				CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		return SellerEntity{
			UserID:   user.UserID,
			Name:     user.Name,
			Email:    user.Email,
			Products: products,
		}
	}

	return SellerEntity{
		UserID:   user.UserID,
		Name:     user.Name,
		Email:    user.Email,
		Products: "never made a product",
	}
}
