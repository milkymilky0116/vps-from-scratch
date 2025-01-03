package server

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/milkymilky0116/vps-from-scratch/internal/repository"
	"github.com/pressly/goose/v3"
)

type Server struct {
	Routes     *http.ServeMux
	Repository *repository.Queries
	DB         *pgxpool.Pool
}

func NewServer(ctx context.Context) *Server {
	conn, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Fail to get pool: %v", err)
	}
	repo := repository.New(conn)
	srv := &Server{
		DB:         conn,
		Repository: repo,
		Routes:     http.NewServeMux(),
	}
	srv.InitRoutes()
	return srv
}

func (s *Server) Run(ctx context.Context, schemas embed.FS) error {
	db := stdlib.OpenDBFromPool(s.DB)
	if err := s.DB.Ping(ctx); err != nil {
		log.Fatalf("Fail to ping to db: %v", err)
	}

	goose.SetBaseFS(schemas)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Fail to set dialect: %v", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Fail to migrate db: %v", err)
	}
	srv := http.Server{
		Addr:    ":8080",
		Handler: s.Routes,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Fail to serve: %v", err)
	}
	return nil
}
