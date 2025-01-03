package server

import "github.com/milkymilky0116/vps-from-scratch/internal/handler"

func (s *Server) InitRoutes() {
	healthCheckRouter := handler.NewHealthCheckHandler()
	repositoryRouter := handler.NewTodoHandler(s.Repository)
	s.Routes.HandleFunc("GET /health_check", healthCheckRouter.HealthCheck)
	s.Routes.HandleFunc("GET /todo/{id}", repositoryRouter.GetTodoById)
	s.Routes.HandleFunc("POST /todo", repositoryRouter.CreateTodo)
}
