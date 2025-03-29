package entity

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
}
