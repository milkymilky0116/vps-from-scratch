package main

import (
	"context"
	"embed"
	"log"

	"github.com/milkymilky0116/vps-from-scratch/internal/server"
)

//go:embed migrations/*.sql
var schemas embed.FS

func main() {
	srv := server.NewServer(context.Background())
	if err := srv.Run(context.Background(), schemas); err != nil {
		log.Fatalf("Fail to launch server : %v", err)
	}
}
