package dbrepo

import (
	"backend/internal/models"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	var movies []*models.Movie
	return movies, nil
}
