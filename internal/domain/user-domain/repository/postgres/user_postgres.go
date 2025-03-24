package postgres

import (
	"database/sql"
	"go-microservice-template/internal/domain/user-domain/entity"

	"github.com/google/uuid"
)

type UserPostgresRepo struct {
	DB *sql.DB
}

// NewUserRepository initializes the repository with a PostgreSQL connection.
func NewUserRepository(db *sql.DB) *UserPostgresRepo {
	return &UserPostgresRepo{DB: db}
}

func (u *UserPostgresRepo) GetAllUsers() ([]entity.User, error) {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// create user in database
func (u *UserPostgresRepo) CreateUser(user entity.User) error {

	_, err := u.DB.Exec("INSERT INTO users (username, email) VALUES ($1,$2)", user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// update user in database
func (u *UserPostgresRepo) UpdateUser(user entity.User) error {
	_, err := u.DB.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func (u *UserPostgresRepo) GetUserByID(id uuid.UUID) (entity.User, error) {
	user := entity.User{}
	err := u.DB.QueryRow("SELECT id,username,email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// delete user by id
func (u *UserPostgresRepo) DeleteUser(id uuid.UUID) error {
	_, err := u.DB.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
