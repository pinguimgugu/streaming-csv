package routes

import (
	"github.com/pinguimgugu/streaming-csv/application/resource"
	"github.com/pinguimgugu/streaming-csv/domain/contract"

	"github.com/labstack/echo"
)

// UserRoutes struct
type UserRoutes struct {
	echo           *echo.Echo
	userRepository contract.UserRepository
}

//NewUserRoutes func
func NewUserRoutes(e *echo.Echo, userRepository contract.UserRepository) *UserRoutes {
	return &UserRoutes{e, userRepository}
}

// Handler
func (a *UserRoutes) Handler() {
	a.echo.GET("/users/csv/", func(c echo.Context) error {
		action := resource.StreamingGetUser{UserRepo: a.userRepository}
		return action.Handler(c)
	})
}
