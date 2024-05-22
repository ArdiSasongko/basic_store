package productcontroller

import "github.com/labstack/echo/v4"

type ProductControllerInterface interface {
	GetProduct(c echo.Context) error
	GetAllProduct(c echo.Context) error
}
