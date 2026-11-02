package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"quiz/internal/models"
	"quiz/internal/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	address := "postgres://postgres:qqquiz@127.0.0.1:5432/quiz_db?sslmode=disable"

	dbUsers, err := sql.Open("pgx", address)

	if err != nil {
		log.Fatalf("ERROR: couldn't launch database connection \t %v\n", err)
	}
	defer dbUsers.Close()

	if err := dbUsers.Ping(); err != nil {
		log.Fatalf("ERROR: couldn't establish a connection \t %v\n", err)
	}

	repoUsers := repository.NewRepository(dbUsers)

	ctx := context.Background()

	user := models.User{
		Login:    "test",
		Password: "test",
		Email:    "test@gmail.com",
	}

	err = repoUsers.CreateNewUser(ctx, &user)

	if err != nil {
		log.Printf("ERROR: couldn't create a new user \t %v\n", err)
	}

	fmt.Printf("Успешно создан пользователь! ID: %d, Время регистрации: %v\n", user.Id, user.CreatedAt)
}
