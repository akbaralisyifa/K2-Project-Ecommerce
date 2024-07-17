package repository

import (
	"ecommerce/internal/features/products"

	"gorm.io/gorm"
)

type ProductModels struct {
	db *gorm.DB
}

func NewProductModels(connect *gorm.DB) products.Query {
	return &ProductModels{
		db: connect,
	}
}

// Add Product
func (pm *ProductModels) AddProduct(newProducts products.Product, userID uint) error {
	cnvData := ToProductsQuery(newProducts)
	cnvData.UserID = userID
	err := pm.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

// Get one data product
func (pm *ProductModels) GetProduct(ID uint) (products.Product, error) {
	var result Products
	err := pm.db.Where("id = ?", ID).First(&result).Error

	if err != nil {
		return products.Product{}, err
	}

	return result.ToProductsEntity(), nil
}

// update data product
func (pm *ProductModels) UpdateProduct(ID uint, updateProduct products.Product) error {
	cnvQuery := ToProductsQuery(updateProduct);

	qry := pm.db.Where("id = ?", ID).Updates(&cnvQuery)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pm *ProductModels) DeleteProduct(ID uint) error {
	qry := pm.db.Where("id = ?", ID).Delete(&Products{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pm *ProductModels) GetAllProducts() ([]products.Product, error) {
	var result []Products
	var resultConvert []products.Product
	err := pm.db.Find(&result).Error
	if err != nil {
		return []products.Product{}, err
	}
	for _, v := range result {
		resultConvert = append(resultConvert, v.ToProductsEntity())
	}
	return resultConvert, nil
}
