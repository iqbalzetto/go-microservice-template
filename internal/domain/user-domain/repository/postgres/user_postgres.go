package postgres

import (
	"database/sql"
	"go-microservice-template/internal/domain/user-domain/entity"
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
