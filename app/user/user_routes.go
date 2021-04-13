package user

import (
	"github.com/labstack/echo/v4"
	"github.com/renkha/go-try/database"
	"github.com/renkha/go-try/helper"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := database.GetDbInstance()
	db.AutoMigrate(User{})
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.UserRegistration,
		},
	}
}
