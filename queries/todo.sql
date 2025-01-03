-- name: FindTodoById :one
SELECT * FROM todo WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todo (context) VALUES ($1) RETURNING *;
