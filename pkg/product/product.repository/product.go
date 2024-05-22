package productrepository

import (
	"basic-store/db/model/domain"
	"errors"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

// GetProduct by id
func (pR *ProductRepo) GetProduct(id int) (domain.Products, error) {
	var product domain.Products

	if err := pR.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

// GetAllProduct get all product
func (pR *ProductRepo) GetAllProduct() ([]domain.Products, error) {
	var products []domain.Products

	if err := pR.DB.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

// Update product
func (pR *ProductRepo) Update(product domain.Products) (domain.Products, error) {
	if err := pR.DB.Model(domain.Products{}).Where("product_id = ?", product.ProductID).Updates(&product).Error; err != nil {
		return product, errors.New("failed to update product")
	}

	return product, nil
}
