package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/server"
)

func main() {
	dsn := "doadmin:BOP@test24New@tcp(localhost:3306)/defaultdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create the new server: %s", err)
	}

	go func() {
		if err := server.StartServer(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.GracefullyShutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
