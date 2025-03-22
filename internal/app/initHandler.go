package app

import (
	"database/sql"
	"go-microservice-template/internal/domain/user-domain/handler"
	"go-microservice-template/internal/domain/user-domain/repository/postgres"
	"go-microservice-template/internal/domain/user-domain/usecase"
)

// TEST COMMIT
// Initialize repositories, usecases and handlers for reuse
func InitUserDomainHandler(db *sql.DB) *handler.UserDomainHandlers {
	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	// Initialize usecases
	userUsecase := usecase.NewUserUsecase(userRepo)
	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)

	return &handler.UserDomainHandlers{
		UserHandler: userHandler,
	}
}
