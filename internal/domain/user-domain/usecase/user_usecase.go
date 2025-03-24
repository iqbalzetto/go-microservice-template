package usecase

import (
	"context"
	"go-microservice-template/internal/domain/user-domain/entity"
	"go-microservice-template/internal/domain/user-domain/repository"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo repository.UserRepo
}

func NewUserUsecase(userRepo repository.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserUsecase) CreateUser(ctx context.Context, user entity.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user entity.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return u.userRepo.DeleteUser(id)
}
