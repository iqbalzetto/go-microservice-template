package repository

import (
	"go-microservice-template/internal/domain/user-domain/entity"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	GetUserByID(id uuid.UUID) (entity.User, error)
	DeleteUser(id uuid.UUID) error
}
