package userservice

import (
	"basic-store/db/model/entity"
	"basic-store/db/model/web"
	"basic-store/helper"
)

type UserServiceInterface interface {
	RegisterUser(user web.UserRequest) (helper.CustomResponse, error)
	RegisterSeller(user web.UserRequest) (helper.CustomResponse, error)
	LoginUser(email, password string) (helper.CustomResponse, error)
	GetHistory(userID int) (entity.BuyerEntity, error)
	GetProduct(userID int) (entity.SellerEntity, error)
}
