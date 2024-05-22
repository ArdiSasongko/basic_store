package orderservice

import (
	"basic-store/db/model/web"
	"basic-store/helper"
)

type OrderServiceInterface interface {
	Create(id int, userid int, order web.OrderRequest) (helper.CustomResponse, error)
}
