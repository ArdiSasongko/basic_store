package productservice

import "basic-store/db/model/entity"

type ProductServiceInterface interface {
	GetProduct(id int) (entity.ProductEntity, error)
	GetAllProduct() ([]entity.ProductEntity, error)
}
