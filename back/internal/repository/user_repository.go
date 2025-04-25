package repository

import (
	"Flowers-store/internal/model"
	"database/sql"
	"errors"
)

// Repository реализует интерфейс для хранения данных
type Repository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (p *postgresRepository) Create(user *model.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := p.db.Exec(query, user.Username, user.Password)
	return err
}

func (p *postgresRepository) GetByUsername(username string) (*model.User, error) {
	var u model.User
	row := p.db.QueryRow(`SELECT id, username, password FROM users WHERE username=$1`,
		username)
	err := row.Scan(&u.ID, &u.Username, &u.Password)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil // Пользователь не найден
	}
	return &u, err
}
