package handler

import (
	"ecommerce/internal/features/products"
	"ecommerce/internal/helpers"
	"ecommerce/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	srv products.Service
}

func NewProductController(s products.Service) products.Handler {
	return &ProductController{
		srv: s,
	}
}

func (uc *ProductController) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ProductInput

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = uc.srv.AddProduct(ToModelProduct(input))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "register success", nil))
	}
}

func (uc *ProductController) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		result, err := uc.srv.GetProduct(uint(ID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "Product successfully retrieved", ToGetProductResponse(result)))
	}
}

func (uc *ProductController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var input ProductInput
		err := c.Bind(&input)

		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "Invalid request parameters", nil))
		}

		err = uc.srv.UpdateProduct(uint(ID), ToModelProduct(input))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Product profile updated", nil))
	}
}

func (uc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		err := uc.srv.DeleteProduct(uint(ID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "Product account deleted", nil))
	}
}
