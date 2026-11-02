package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"quiz/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateNewUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (login, password, email)
		VALUES ($1, $2, $3)
		RETURNING id, time_created_at
	`
	err := repo.db.QueryRowContext(
		ctx,
		query,
		user.Login,
		user.Password,
		user.Email,
	).Scan(&user.Id, &user.CreatedAt)

	if err != nil {
		log.Print(err)
		return errors.New("failed to create user")
	}

	return nil

}
