// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: todo.sql

package repository

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todo (context) VALUES ($1) RETURNING id, context, created_at, updated_at
`

func (q *Queries) CreateTodo(ctx context.Context, argContext string) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, argContext)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Context,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findTodoById = `-- name: FindTodoById :one
SELECT id, context, created_at, updated_at FROM todo WHERE id = $1
`

func (q *Queries) FindTodoById(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRow(ctx, findTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Context,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}