package handler

import (
	"ecommerce/internal/features/users"
	"ecommerce/internal/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserContorler struct {
	srv users.Service
}

func NewUserController(s users.Service) users.Handler{
	return &UserContorler{
		srv: s,
	}
};

func (uc *UserContorler) Register() echo.HandlerFunc {
	return func(c echo.Context) error{
		var input UserRequest;

		err := c.Bind(&input);
		if err != nil {
			log.Println("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		};

		err = uc.srv.Register(ToModelUser(input));
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "register success", nil))
	}
}

func (uc *UserContorler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest;

		err := c.Bind(&input);
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}

		_, token, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "success", "data": ToLoginResponse(token)})
	}
}


