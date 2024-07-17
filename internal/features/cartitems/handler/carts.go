package handler

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/helpers"
	"ecommerce/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CartItemController struct {
	srv cartitems.Service
}

func NewCartItemController(s cartitems.Service) cartitems.Handler {
	return &CartItemController{
		srv: s,
	}
}

func (uc *CartItemController) AddCartItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		var input CartItemInput

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = uc.srv.AddCartItem(ToModelCartItem(input), uint(userID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Cart item successfully added", nil))
	}
}

func (uc *CartItemController) DeleteCartItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		CartID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Invalid request parameters", nil))
		}

		err = uc.srv.DeleteCartItem(uint(CartID), uint(userID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "Cart deleted", nil))
	}
}

func (uc *CartItemController) GetAllCartItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))
		responseData, err := uc.srv.GetAllCartItems(uint(userID))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "not found") {
				errCode = 400
			}
			return c.JSON(errCode, helpers.ResponseFormat(errCode, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "All Carts retrieved", responseData))
	}
}
