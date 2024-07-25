package service

import (
	"ecommerce/internal/features/products"
	"ecommerce/internal/utils"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ProductServices struct {
	qry products.Query
	pwd utils.HashingPwInterface
	vld utils.ValidatorUtilityInterface
	jwt utils.JwtUtilityInterface
	exl utils.ExcelUtilityInterface
}

func NewProductService(q products.Query, p utils.HashingPwInterface, v utils.ValidatorUtilityInterface, j utils.JwtUtilityInterface, x utils.ExcelUtilityInterface) products.Service {
	return &ProductServices{
		qry: q,
		pwd: p,
		vld: v,
		jwt: j,
		exl: x,
	}
}

func (ps *ProductServices) AddProduct(newProduct products.Product, userID uint) error {

	// add product
	err := ps.qry.AddProduct(newProduct, userID)
	if err != nil {
		log.Println("add product sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (ps *ProductServices) GetProduct(ID uint) (products.Product, error) {
	result, err := ps.qry.GetProduct(ID)

	if err != nil {
		log.Print("get Product query error", err.Error())
		return products.Product{}, errors.New("internal server error")
	}

	return result, nil
}

func (ps *ProductServices) UpdateProduct(productID uint, userID uint, updatedProduct products.Product) error {
	result, err := ps.qry.GetProduct(productID)
	if err != nil {
		log.Print("not found", err.Error())
		return errors.New("internal server error")
	}

	if result.UserID != userID {
		return errors.New(" unauthorize update action")
	}

	// update Product
	err = ps.qry.UpdateProduct(productID, updatedProduct)

	if err != nil {
		log.Print("update Product query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}

func (ps *ProductServices) DeleteProduct(productID uint, userID uint) error {
	result, err := ps.qry.GetProduct(productID)
	if err != nil {
		log.Print("not found", err.Error())
		return errors.New("internal server error")
	}

	if result.UserID != userID {
		return errors.New(" unauthorize delete action")
	}
	err = ps.qry.DeleteProduct(productID)

	if err != nil {
		log.Print("delete Product query error", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (ps *ProductServices) GetAllProducts() ([]products.Product, error) {

	result, err := ps.qry.GetAllProducts()
	msg := "internal server error"
	if err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "not found"
		}
		return []products.Product{}, errors.New(msg)
	}

	return result, nil
}

func (ps *ProductServices) GetAllUserProducts(userID uint) ([]products.Product, error) {

	result, err := ps.qry.GetAllUserProducts(userID)
	msg := "internal server error"
	if err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "not found"
		}
		return []products.Product{}, errors.New(msg)
	}

	return result, nil
}
func (ps *ProductServices) GetAllOtherUserProducts(userID uint) ([]products.Product, error) {

	result, err := ps.qry.GetAllOtherUserProducts(userID)
	msg := "internal server error"
	if err != nil {

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "not found"
		}
		return []products.Product{}, errors.New(msg)
	}

	return result, nil
};


func (ps *ProductServices) GenerateProductsExcel() ([]byte, error){
	products, err := ps.qry.GetAllProducts();

	if err != nil {
		return nil,  errors.New("failed to retrieve products")
	};

	excelFile, err := ps.exl.DownloadExcel(products)

	if err != nil {
		fmt.Println(err)
		return nil,  errors.New("failed to download product")
	};

	return excelFile, nil;
}