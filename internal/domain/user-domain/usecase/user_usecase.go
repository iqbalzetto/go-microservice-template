package usecase

import (
	"context"
	"fmt"
	"go-microservice-template/internal/domain/user-domain/dto"
	"go-microservice-template/internal/domain/user-domain/entity"
	"go-microservice-template/internal/domain/user-domain/repository"

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

func (u *UserUsecase) UploadProfilePicture(ctx context.Context, id uuid.UUID, file dto.InputFileDTO) error {
	//upload profile picture to minio
	if u.mc == nil {
		return fmt.Errorf("MinIO client is not initialized")
	}

	_, err := u.mc.PutObject(context.Background(), "user-service", "test/"+file.Name, file.Reader, file.Size, minio.PutObjectOptions{
		ContentType: file.Type,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	//upload profile picture to database

	return nil

}

func (u *UserUsecase) ExportToExcel(ctx context.Context, id uuid.UUID) (string, error) {
	// Get user by ID
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}

	// Create Excel file and write user data
	fileName := fmt.Sprintf("%s.xlsx", user.Username)
	err = createExcelFile(fileName, user)
	if err != nil {
		return "", fmt.Errorf("failed to create Excel file: %v", err)
	}

	return fileName, nil
}
