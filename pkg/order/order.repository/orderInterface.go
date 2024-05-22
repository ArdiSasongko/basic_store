package orderrepository

import "basic-store/db/model/domain"

type OrderRepoInterface interface {
	Create(order domain.Orders) (domain.Orders, error)
}
