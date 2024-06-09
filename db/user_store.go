package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/idomath/StrictlyRecipes/types"
)

type UserStore struct {
	Db *sql.DB
}

func (s *UserStore) InsertUser(user types.User) (int, error) {
	ctx, cancel := context.WithCancel(context.Background(), 3*time.Second)
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
