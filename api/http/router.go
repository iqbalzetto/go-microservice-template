package router

import (
	userRoutes "go-microservice-template/internal/domain/user-domain/handler/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	userRoutes.InitRoutes(e)
}
