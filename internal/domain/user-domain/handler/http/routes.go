package http

import (
	"go-microservice-template/internal/domain/user-domain/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userDomainHandlers *handler.UserDomainHandlers) {

	v1 := e.Group("/v1")

	userRoutes := v1.Group("/users")
	userRoutes.GET("", userDomainHandlers.UserHandler.GetAllUsers)

}
