package http

import (
	"go-microservice-template/internal/domain/user-domain/handler"
	userRoutes "go-microservice-template/internal/domain/user-domain/handler/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userDomainHandlers *handler.UserDomainHandlers) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	userRoutes.InitRoutes(e, userDomainHandlers)
}
