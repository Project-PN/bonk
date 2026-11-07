package model

import "github.com/jackc/pgx/v5/pgxpool"

// repository for a model
type Repository[T any] struct {
	pool *pgxpool.Pool
}

// create a model repository
func New[T any](pool *pgxpool.Pool) *Repository[T] {
	return &Repository[T]{pool: pool}
}
