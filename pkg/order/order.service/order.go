package orderservice

import (
	"basic-store/db/model/domain"
	"basic-store/db/model/web"
	"basic-store/helper"
	orderrepository "basic-store/pkg/order/order.repository"
	productrepository "basic-store/pkg/product/product.repository"
	"errors"
)

type OrderService struct {
	Repo        orderrepository.OrderRepoInterface
	ProductRepo productrepository.ProductRepoInterface
}

func NewOrderService(repo orderrepository.OrderRepoInterface, productRepo productrepository.ProductRepoInterface) *OrderService {
	return &OrderService{
		Repo:        repo,
		ProductRepo: productRepo,
	}
}

// Create order
func (oS *OrderService) Create(id int, userid int, order web.OrderRequest) (helper.CustomResponse, error) {
	product, err := oS.ProductRepo.GetProduct(id)

	if err != nil {
		return nil, err
	}

	totalPrice := product.Price * order.Quantity

	if order.TotalPrice < totalPrice {
		return nil, errors.New("total price is not enough")
	}

	newQuantity := product.Quantity - order.Quantity

	if product.Quantity < order.Quantity {
		return nil, errors.New("quantity is not enough")
	}

	newProduct := domain.Products{
		ProductID:  id,
		SellerIDFK: product.SellerIDFK,
		Name:       product.Name,
		Quantity:   newQuantity,
		Price:      product.Price,
	}

	_, err = oS.ProductRepo.Update(newProduct)

	if err != nil {
		return nil, err
	}

	newOrder := domain.Orders{
		UserIDFK:    userid,
		ProductIDFK: id,
		Name:        product.Name,
		Quantity:    order.Quantity,
		TotalPrice:  totalPrice,
	}

	result, err := oS.Repo.Create(newOrder)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"order": result,
	}

	return data, nil
}
