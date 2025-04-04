package handler

import (
	"database/sql"
	"errors"
	"go-microservice-template/internal/domain/user-domain/dto"
	"go-microservice-template/internal/domain/user-domain/entity"
	"go-microservice-template/internal/domain/user-domain/usecase"
	res "go-microservice-template/pkg/response"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {

	user, err := h.userUsecase.GetAllUsers(c.Request().Context())
	if err != nil {
		log.Println(err)
		return res.JSON(c, res.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if user == nil {
		return res.JSON(c, res.StatusNotFound, nil)
	}

	return res.JSON(c, res.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.userUsecase.CreateUser(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.userUsecase.GetUserByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user entity.User

	// Parse user ID from the request URL (e.g., /users/:id)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Check if user exists before updating
	existingUser, err := h.userUsecase.GetUserByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Handle case where user is not found
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user.ID = existingUser.ID // Assign the parsed ID to user struct

	if err := h.userUsecase.UpdateUser(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := h.userUsecase.DeleteUser(c.Request().Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (h *UserHandler) UploadProfilePicture(c echo.Context) error {
	// Parse user ID from the request URL (e.g., /users/:id)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "file not found")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open file")
	}
	defer src.Close()

	// Read the first 512 bytes to detect the content type
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return err
	}

	// Reset the file pointer back to the beginning for future reads
	src.Seek(0, 0)
	contentType := http.DetectContentType(buffer)

	// create file DTO
	inputFileDTO := dto.InputFileDTO{
		Name:      fileHeader.Filename,
		Size:      fileHeader.Size,
		Type:      contentType,
		Extension: fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:],
		Reader:    src,
	}

	if err := h.userUsecase.UploadProfilePicture(c.Request().Context(), id, inputFileDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Profile picture updated successfully"})
}
