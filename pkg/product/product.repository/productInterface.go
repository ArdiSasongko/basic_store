package productrepository

import "basic-store/db/model/domain"

type ProductRepoInterface interface {
	GetProduct(id int) (domain.Products, error)
	GetAllProduct() ([]domain.Products, error)
	Update(product domain.Products) (domain.Products, error)
}
