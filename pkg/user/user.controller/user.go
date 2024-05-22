package usercontroller

import (
	"basic-store/db/model/web"
	"basic-store/helper"
	userservice "basic-store/pkg/user/user.service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service userservice.UserServiceInterface
}

func NewUserController(service userservice.UserServiceInterface) *UserController {
	return &UserController{
		service: service,
	}
}

func (uC *UserController) RegisterUser(c echo.Context) error {
	newUser := new(web.UserRequest)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	result, err := uC.service.RegisterUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success Register", result))
}

func (uC *UserController) RegisterSeller(c echo.Context) error {
	newUser := new(web.UserRequest)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	result, err := uC.service.RegisterSeller(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success Register as Seller", result))
}

func (uC *UserController) LoginUser(c echo.Context) error {
	loginUser := new(web.UserLoginRequest)

	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(loginUser); err != nil {
		return err
	}

	login, errLogin := uC.service.LoginUser(loginUser.Email, loginUser.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Login", login))
}

func (uC *UserController) GetHistory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(*helper.CustomClaims)

	result, err := uC.service.GetHistory(claims.UserID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get History", result))
}

func (uC *UserController) GetProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(*helper.CustomClaims)

	result, err := uC.service.GetProduct(claims.UserID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success Get Product", result))
}
