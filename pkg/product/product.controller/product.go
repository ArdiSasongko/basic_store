package productcontroller

import (
	"basic-store/helper"
	productservice "basic-store/pkg/product/product.service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Service productservice.ProductServiceInterface
}

func NewProductController(service productservice.ProductServiceInterface) *ProductController {
	return &ProductController{
		Service: service,
	}
}

// GetProduct by id
func (pC *ProductController) GetProduct(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errId.Error(), nil))
	}

	result, err := pC.Service.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "found product", result))
}

// GetAllProduct get all product
func (pC *ProductController) GetAllProduct(c echo.Context) error {
	result, err := pC.Service.GetAllProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "found all product", result))
}
