package usecase

import (
	"context"
	"fmt"
	"go-microservice-template/internal/domain/user-domain/entity"
	"go-microservice-template/internal/domain/user-domain/repository"
	"os"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type UserUsecase struct {
	userRepo repository.UserRepo
	mc       *minio.Client
}

func NewUserUsecase(
	userRepo repository.UserRepo,
	mc *minio.Client,
) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		mc:       mc,
	}
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

func (u *UserUsecase) UploadProfilePicture(ctx context.Context) error {
	//upload profile picture to minio
	if u.mc == nil {
		return fmt.Errorf("MinIO client is not initialized")
	}

	file, err := os.Open("./apple.jpeg")
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	_, err = u.mc.PutObject(context.Background(), "user-service", "testing", file, fileStat.Size(), minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	return nil

}
