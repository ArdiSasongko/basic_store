package usercontroller

import "github.com/labstack/echo/v4"

type UserControllerInterface interface {
	RegisterUser(c echo.Context) error
	RegisterSeller(c echo.Context) error
	LoginUser(c echo.Context) error
	GetHistory(c echo.Context) error
	GetProduct(c echo.Context) error
}
