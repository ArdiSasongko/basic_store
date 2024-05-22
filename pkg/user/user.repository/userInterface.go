package userrepository

import "basic-store/db/model/domain"

type UserRepoInterface interface {
	RegisterUser(user domain.Users) (domain.Users, error)
	RegisterSeller(user domain.Users) (domain.Users, error)
	GetEmail(email string) (domain.Users, error)
	GetHistory(userID int) (domain.Users, error)
	GetProduct(userID int) (domain.Users, error)
}
