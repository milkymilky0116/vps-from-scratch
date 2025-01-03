package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/milkymilky0116/vps-from-scratch/internal/dto"
	"github.com/milkymilky0116/vps-from-scratch/internal/repository"
	testutil "github.com/milkymilky0116/vps-from-scratch/internal/test_util"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {
	dbUrl, err := testutil.LaunchPostgresContainer()

	assert.NoError(t, err)

	conn, err := pgxpool.New(context.Background(), *dbUrl)
	assert.NoError(t, err)
	db := stdlib.OpenDBFromPool(conn)
	if err := goose.SetDialect("postgres"); err != nil {
		t.Fatalf("Fail to set dialect: %v", err)
	}

	if err := goose.Up(db, "../../migrations"); err != nil {
		t.Fatalf("Fail to migrate db: %v", err)
	}
	repo := repository.New(conn)
	todoHandler := NewTodoHandler(repo)

	reqBody := dto.CreateTodoRequest{
		Context: "test",
	}
	data, err := json.Marshal(reqBody)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	todoHandler.CreateTodo(writer, req)

	var result repository.Todo
	resp := writer.Result()
	resp.Header.Set("Content-Type", "application/json")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(body, &result)

	assert.NoError(t, err)

	assert.Equal(t, result.Context, "test")
}
