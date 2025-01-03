-- +goose Up
CREATE TABLE test (
  id BIGSERIAL PRIMARY KEY,
  test varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
DROP TABLE test;
