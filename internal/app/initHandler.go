package app

import (
	"database/sql"
	"go-microservice-template/internal/domain/user-domain/handler"
	"go-microservice-template/internal/domain/user-domain/repository/postgres"
	"go-microservice-template/internal/domain/user-domain/usecase"

	"github.com/minio/minio-go/v7"
)

// TEST COMMIT
// Initialize repositories, usecases and handlers for reuse
func InitUserDomainHandler(db *sql.DB, mc *minio.Client) *handler.UserDomainHandlers {
	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	// Initialize usecases
	userUsecase := usecase.NewUserUsecase(userRepo, mc)
	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)

	return &handler.UserDomainHandlers{
		UserHandler: userHandler,
	}
}
