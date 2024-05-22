package userrepository

import (
	"basic-store/db/model/domain"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (uR *UserRepo) RegisterUser(user domain.Users) (domain.Users, error) {
	if err := uR.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Users{}, errors.New("email already exists")
		}
		return domain.Users{}, errors.New("failed to register user")
	}

	return user, nil
}

func (uR *UserRepo) RegisterSeller(user domain.Users) (domain.Users, error) {
	if err := uR.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.Users{}, errors.New("email already exists")
		}
		return domain.Users{}, errors.New("failed to register seller")
	}

	return user, nil
}

func (uR *UserRepo) GetEmail(email string) (domain.Users, error) {
	var user domain.Users
	if err := uR.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return domain.Users{}, errors.New("email not found")
	}

	return user, nil
}

func (uR *UserRepo) GetHistory(userID int) (domain.Users, error) {
	var User domain.Users
	if err := uR.DB.Preload("Orders").First(&User, "user_id = ?", userID).Error; err != nil {
		return domain.Users{}, errors.New("user not found")
	}
	return User, nil
}

func (uR *UserRepo) GetProduct(userID int) (domain.Users, error) {
	var User domain.Users
	if err := uR.DB.Preload("Products").First(&User, "user_id = ?", userID).Error; err != nil {
		return domain.Users{}, errors.New("user not found")
	}
	return User, nil
}
