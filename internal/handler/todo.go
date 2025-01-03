package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/milkymilky0116/vps-from-scratch/internal/dto"
	"github.com/milkymilky0116/vps-from-scratch/internal/repository"
)

type TodoHandler struct {
	Repository *repository.Queries
}

func (t *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo, err := t.Repository.FindTodoById(context.Background(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if todo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var reqBody dto.CreateTodoRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	todo, err := t.Repository.CreateTodo(context.Background(), reqBody.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewTodoHandler(repo *repository.Queries) *TodoHandler {
	return &TodoHandler{
		Repository: repo,
	}
}
