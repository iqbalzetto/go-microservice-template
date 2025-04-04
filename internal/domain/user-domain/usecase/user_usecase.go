package usecase

import (
	"bytes"
	"context"
	"fmt"
	"go-microservice-template/internal/domain/user-domain/dto"
	"go-microservice-template/internal/domain/user-domain/entity"
	"go-microservice-template/internal/domain/user-domain/repository"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/xuri/excelize/v2"
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

func (u *UserUsecase) ExportToExcel(ctx context.Context) (*bytes.Buffer, error) {
	// Get all users
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	// Create a new Excel file
	f := excelize.NewFile()
	// Create a new sheet
	sheetName := "Users"
	f.SetSheetName("Sheet1", sheetName)

	// Set the headers
	headers := []string{"ID", "Username", "Email"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Users", cell, header)
	}
	// Set the data
	for i, user := range users {
		cellID, _ := excelize.CoordinatesToCellName(1, i+2)
		cellUsername, _ := excelize.CoordinatesToCellName(2, i+2)
		cellEmail, _ := excelize.CoordinatesToCellName(3, i+2)
		f.SetCellValue("Users", cellID, user.ID)
		f.SetCellValue("Users", cellUsername, user.Username)
		f.SetCellValue("Users", cellEmail, user.Email)
	}
	// Save to buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return &buf, nil
}
