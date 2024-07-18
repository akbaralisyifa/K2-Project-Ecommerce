package handler

import (
	"ecommerce/internal/features/products"
	"ecommerce/internal/helpers"
	"ecommerce/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		var input ProductInput

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		file, err := c.FormFile("product_image")
		if err == nil {
			// open file
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			defer src.Close()

			urlImage, err := utils.UploadToCloudinary(src, file.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			input.ImageUrl = urlImage
		}

		err = uc.srv.AddProduct(ToModelProduct(input), uint(userID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Product successfully created", nil))
	}
}

func (uc *ProductController) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}
		result, err := uc.srv.GetProduct(uint(productID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "Product successfully retrieved", result))
	}
}

func (uc *ProductController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}
		var input ProductInput
		err = c.Bind(&input)

		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "Invalid request parameters", nil))
		}

		file, err := c.FormFile("product_image")
		if err == nil {
			// open file
			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			defer src.Close()

			urlImage, err := utils.UploadToCloudinary(src, file.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "profile image error", nil))
			}
			input.ImageUrl = urlImage
		}

		err = uc.srv.UpdateProduct(uint(productID), uint(userID), ToModelProduct(input))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Product updated", nil))
	}
}

func (uc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		productID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		err = uc.srv.DeleteProduct(uint(productID), uint(userID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "Product deleted", nil))
	}
}

func (uc *ProductController) GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		responseData, err := uc.srv.GetAllProducts()
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "not found") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "All products retrieved", responseData))
	}
}

func (uc *ProductController) GetAllProductsByOwner() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		isPersonal, err := strconv.ParseBool(c.QueryParam("isPersonal"))
		if err == nil {
			if isPersonal {
				responseData, err := uc.srv.GetAllUserProducts(uint(userID))
				if err != nil {
					errCode := 500
					if strings.ContainsAny(err.Error(), "not found") {
						errCode = 400
					}
					return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
				}
				return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "All products retrieved", responseData))
			} else {
				responseData, err := uc.srv.GetAllOtherUserProducts(uint(userID))
				if err != nil {
					errCode := 500
					if strings.ContainsAny(err.Error(), "not found") {
						errCode = 400
					}
					return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
				}
				return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "All products retrieved", responseData))
			}
		}

		responseData, err := uc.srv.GetAllProducts()
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "not found") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "All products retrieved", responseData))
	}
}
