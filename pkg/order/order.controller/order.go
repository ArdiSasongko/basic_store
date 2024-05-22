package ordercontroller

import (
	"basic-store/db/model/web"
	"basic-store/helper"
	orderservice "basic-store/pkg/order/order.service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Service orderservice.OrderServiceInterface
}

func NewOrderController(service orderservice.OrderServiceInterface) *OrderController {
	return &OrderController{
		Service: service,
	}
}

// Create order
func (oC *OrderController) Create(c echo.Context) error {
	newOrder := new(web.OrderRequest)
	token := c.Get("user").(*jwt.Token)
	claims, _ := token.Claims.(*helper.CustomClaims)
	userId := claims.UserID

	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errId.Error(), nil))
	}

	if err := c.Bind(newOrder); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newOrder); err != nil {
		return err
	}

	result, err := oC.Service.Create(id, userId, *newOrder)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "success buy", result))
}
