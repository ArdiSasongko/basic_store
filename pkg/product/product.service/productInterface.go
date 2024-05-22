package productservice

import (
	"basic-store/db/model/entity"
	productrepository "basic-store/pkg/product/product.repository"
)

type ProductService struct {
	Repo productrepository.ProductRepoInterface
}

func NewProductService(repo productrepository.ProductRepoInterface) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

// GetProduct by id
func (pS *ProductService) GetProduct(id int) (entity.ProductEntity, error) {
	result, err := pS.Repo.GetProduct(id)

	if err != nil {
		return entity.ProductEntity{}, err
	}

	return entity.ToProduct(result), nil
}

// GetAllProduct get all product
func (pS *ProductService) GetAllProduct() ([]entity.ProductEntity, error) {
	result, err := pS.Repo.GetAllProduct()

	if err != nil {
		return []entity.ProductEntity{}, err
	}

	return entity.ToProductList(result), nil
}
