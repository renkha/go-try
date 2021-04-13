package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/renkha/go-try/helper"
)

type userHandler struct {
	userService Services
}

func NewHandler(userService Services) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) UserRegistration(c echo.Context) error {
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errors := helper.ErrorFormatter(err)
		errMessage := helper.M{"errors": errors}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "failed", errMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	existEmail := h.userService.CheckExistEmail(*req)
	if existEmail != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", existEmail.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.CreateUser(*req)
	if err != nil {
		errors := helper.ErrorFormatter(err)
		errMessage := helper.M{"errors": errors}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "failed", errMessage)
		return c.JSON(http.StatusBadRequest, response)
	}
	userData := UserResponseFormatter(newUser)

	response := helper.ResponseFormatter(http.StatusOK, "success", "succes user", userData)

	return c.JSON(http.StatusOK, response)
}
