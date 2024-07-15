package service

import (
	"ecommerce/internal/features/products"
	"ecommerce/internal/utils"
	"errors"
	"log"
)

type ProductServices struct {
	qry products.Query
	pwd utils.HashingPwInterface
	vld utils.ValidatorUtilityInterface
	jwt utils.JwtUtilityInterface
}

func NewProductService(q products.Query, p utils.HashingPwInterface, v utils.ValidatorUtilityInterface, j utils.JwtUtilityInterface) products.Service {
	return &ProductServices{
		qry: q,
		pwd: p,
		vld: v,
		jwt: j,
	}
}

func (ps *ProductServices) AddProduct(newProduct products.Product) error {
	// validasi data
	err := ps.vld.AddProductValidation(newProduct.Name, newProduct.Price, newProduct.Stock)
	if err != nil {
		log.Println("validator add product error", err.Error())
		return err
	}

	// add product
	err = ps.qry.AddProduct(newProduct)
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

func (ps *ProductServices) UpdateProduct(ID uint, updateProduct products.Product) error {

	// update Product
	err := ps.qry.UpdateProduct(ID, updateProduct)

	if err != nil {
		log.Print("update Product query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}

func (ps *ProductServices) DeleteProduct(ID uint) error {
	err := ps.qry.DeleteProduct(ID)

	if err != nil {
		log.Print("delete Product query error", err.Error())
		return errors.New("interval server error")
	}

	return nil
}
