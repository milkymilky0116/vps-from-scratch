package testutil

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func LaunchPostgresContainer() (*string, error) {
	postgresContainer, err := postgres.Run(context.Background(),
		"docker.io/postgres:latest",
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)))
	if err != nil {
		return nil, err
	}
	connectionString, err := postgresContainer.ConnectionString(context.Background())
	if err != nil {
		return nil, err
	}
	return &connectionString, nil
}
