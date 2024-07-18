package handler

import (
	"ecommerce/internal/features/cartitems"
	"ecommerce/internal/features/orders"
	"ecommerce/internal/helpers"
	"ecommerce/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	srv  orders.Service
	sCrt cartitems.Service
}

func NewOrderController(s orders.Service, c cartitems.Service) orders.Handler {
	return &OrderController{
		srv:  s,
		sCrt: c,
	}
}

func (oc *OrderController) GetAllOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		result, err := oc.srv.GetAllOrder(uint(userID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Orders retrieved", ToOrderResponse(result)))
	}
}

func (oc *OrderController) GetAllOrderHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		result, err := oc.srv.GetAllOrderHistory(uint(userID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server errror", nil))
		}
		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "Orders retrieved", ToOrderResponse(result)))
	}
}

func (oc *OrderController) Checkout() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var orderInput CreateOrderRequest
		err := c.Bind(&orderInput)
		if err != nil {
			return err
		}
		newOrder := ToModelOrders(orderInput, uint(userID), "pending")

		cartItems, _, err := oc.sCrt.GetAllCartItems(uint(userID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "get all cart error", nil))
		}

		respond, err := oc.srv.Checkout(uint(userID), newOrder, cartItems)

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "checkout success", ToCheckoutResponse(respond)))
	}
}

func (oc *OrderController) UpdateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderUpdate MidTransRequest
		err := c.Bind(&orderUpdate)
		if err != nil {
			return err
		}
		orderID, _ := strconv.Atoi(orderUpdate.OrderID)
		newOrderUpdate := ToModelOrders2(uint(orderID), orderUpdate)
		err = oc.srv.UpdateOrder(newOrderUpdate.ID, newOrderUpdate)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "update order error", nil))
		}
		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "payment success", nil))
	}
}
