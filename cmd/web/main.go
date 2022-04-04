package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ischenkx/vk-test-task/cmd/web/config"
	"github.com/ischenkx/vk-test-task/internal/app"
	"github.com/ischenkx/vk-test-task/internal/impl/authorizer/jwtauth"
	"github.com/ischenkx/vk-test-task/internal/impl/data/postgres"
	"github.com/ischenkx/vk-test-task/internal/impl/events/evbus"
	"github.com/ischenkx/vk-test-task/internal/transport/web"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"time"
)

func getConfigFileName() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("config filename not provided")
	}
	return os.Args[1], nil
}

func main() {
	ctx := context.Background()

	configFileName, err := getConfigFileName()

	if err != nil {
		log.Fatalln(err)
		return
	}

	cfg, err := config.FromFile(configFileName)

	if err != nil {
		log.Fatalln(err)
		return
	}

	// "postgres://postgres:123456@127.0.0.1:5432/vk_test"

	pg, err := pgxpool.Connect(ctx, cfg.Postgres.URL)

	if err != nil {
		log.Fatalln("failed to connect to db:", err)
		return
	} else {
		log.Println("connected to db...")
	}

	repo := postgres.NewRepo(pg)

	if err := repo.InitializeTables(ctx); err != nil {
		log.Fatalln("failed to initialize table:", err)
		return
	} else {
		log.Println("initialized tables...")
	}

	auth := jwtauth.New([]byte(cfg.JWT.Key), time.Duration(cfg.JWT.ExpirationTime*1000))

	bus := evbus.NewBus()

	application := app.New(repo, auth, bus)

	mux := web.NewRouter(application)
	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Addr, cfg.HTTP.Port)

	log.Printf("starting http server (%s)...\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalln("failed to start the server")
	}
}
