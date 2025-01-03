-- +goose Up
CREATE TABLE todo (
  id BIGSERIAL PRIMARY KEY,
  context varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE todo;

