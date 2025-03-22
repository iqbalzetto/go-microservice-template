package repository

import "go-microservice-template/internal/domain/user-domain/entity"

type UserRepo interface {
	GetAllUsers() ([]entity.User, error)
}
