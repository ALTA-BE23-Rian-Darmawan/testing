package handler

import (
	users "BE23TODO/features/Users"
	"BE23TODO/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.ServiceUserInterface
}

func New(us users.ServiceUserInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	// Membaca data dari body permintaan
	newUser := UserRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+errBind.Error(), nil))
	}

	// Mapping request ke struct User
	dataUser := users.User{
		FullName:    newUser.FullName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
	}

	// Memanggil service layer untuk menyimpan data
	if errInsert := uh.userService.RegistrasiAccount(dataUser); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Registrasi gagal: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "user registration failed: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "user resgistration succcessful", nil))
}

func (uh *UserHandler) Login(c echo.Context) error {
	// Membaca data dari request body
	loginReq := LoginRequest{}

	if errBind := c.Bind(&loginReq); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+errBind.Error(), nil))
	}

	// Melakukan login
	_, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "user login failed: "+err.Error(), nil))
	}

	// Mengembalikan respons dengan token
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "user login successful", echo.Map{"token": token}))
}
