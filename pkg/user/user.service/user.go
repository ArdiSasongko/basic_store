package userservice

import (
	"basic-store/db/model/domain"
	"basic-store/db/model/entity"
	"basic-store/db/model/web"
	"basic-store/helper"
	userrepository "basic-store/pkg/user/user.repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo  userrepository.UserRepoInterface
	token helper.TokenUseCaseInterface
}

func NewUserService(repo userrepository.UserRepoInterface, token helper.TokenUseCaseInterface) *UserService {
	return &UserService{
		repo:  repo,
		token: token,
	}
}

func (uS *UserService) RegisterUser(user web.UserRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return nil, errHash
	}

	userDomain := domain.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passHash),
		Role:     "buyer",
	}

	result, err := uS.repo.RegisterUser(userDomain)

	if err != nil {
		return nil, err
	}

	saveData := helper.CustomResponse{
		"user_id": result.UserID,
		"name":    result.Name,
		"email":   result.Email,
		"role":    result.Role,
	}

	return saveData, nil
}

func (uS *UserService) RegisterSeller(user web.UserRequest) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return nil, errHash
	}

	userDomain := domain.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(passHash),
		Role:     "seller",
	}

	result, err := uS.repo.RegisterUser(userDomain)

	if err != nil {
		return nil, err
	}

	saveData := helper.CustomResponse{
		"user_id": result.UserID,
		"name":    result.Name,
		"email":   result.Email,
		"role":    result.Role,
	}

	return saveData, nil
}

func (uS *UserService) LoginUser(email, password string) (helper.CustomResponse, error) {
	user, err := uS.repo.GetEmail(email)
	if err != nil {
		return nil, err
	}
	if errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errPass != nil {
		return nil, errors.New("password not match")
	}

	expiredToken := time.Now().Local().Add(5 * time.Minute)
	claims := helper.CustomClaims{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "basic-store",
			ExpiresAt: jwt.NewNumericDate(expiredToken),
		},
	}

	token, errToken := uS.token.GeneraredToken(claims)

	if errToken != nil {
		return nil, errToken
	}

	data := helper.CustomResponse{
		"token": token,
		"exp":   expiredToken,
	}

	return data, nil
}

func (uS *UserService) GetHistory(userID int) (entity.BuyerEntity, error) {
	user, err := uS.repo.GetHistory(userID)

	if err != nil {
		return entity.BuyerEntity{}, err
	}

	return entity.ToBuyerHistory(user), nil
}

func (uS *UserService) GetProduct(userID int) (entity.SellerEntity, error) {
	user, err := uS.repo.GetProduct(userID)

	if err != nil {
		return entity.SellerEntity{}, err
	}

	return entity.ToSellerProduct(user), nil
}
