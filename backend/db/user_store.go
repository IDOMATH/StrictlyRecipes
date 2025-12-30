package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/idomath/StrictlyRecipes/types"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{Db: db}
}

func (s *UserStore) InsertUser(user types.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	statement := `insert into users (email, password_hash)
				  values ($1, $2) returning id`

	err := s.Db.QueryRowContext(ctx, statement, user.email, user.PasswordHash)
	if err != nil {
		return 0, nil
	}
	return newId, nil
}

func (s *UserStore) Authenticate(email, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var passwordHash string
	statement := `select id, password_hash from users where email = $1`

	err := s.Db.QueryRowContext(ctx, statement, email).Scan(&id, &passwordHash)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}
