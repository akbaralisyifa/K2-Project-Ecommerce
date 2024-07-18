package handler

import (
	"ecommerce/internal/features/users"
	"ecommerce/internal/helpers"
	"ecommerce/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv users.Service
}

func NewUserController(s users.Service) users.Handler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UserRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad requeste", nil))
		}

		err = uc.srv.Register(ToModelUser(input))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "register success", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}

		_, token, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "user logged in", ToLoginResponse(token)))
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		result, err := uc.srv.GetUser(uint(ID))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "User successfully retrieved", ToGetUserResponse(result)))
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		var updateUser GetUpdateRequest
		err := c.Bind(&updateUser)
		if err != nil {
			log.Println("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "Invalid request parameters", nil))
		}

		// Bagian updaload Image
		file, err := c.FormFile("image_profile")

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
			updateUser.ImgProfile = urlImage

		}

		err = uc.srv.UpdateUser(uint(ID), ToRequertUser(updateUser))
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "User profile updated", nil))
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := utils.NewJwtUtility().DecodToken(c.Get("user").(*jwt.Token))

		err := uc.srv.DeleteUser(uint(ID))

		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "User account deleted", nil))
	}
}
