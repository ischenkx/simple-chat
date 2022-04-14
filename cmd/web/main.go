package main

import (
	"context"
	"flag"
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
	"time"
)

var EnvConfigFlag = flag.Bool("env-config", false, "tells the program to read the config from environment variables")
var FileConfigFlag = flag.String("file-config", "config.yml", "tells the program to read the config from a file")

func getConfig() (config.Config, error) {
	if *EnvConfigFlag {
		return config.FromENV()
	}

	return config.FromFile(*FileConfigFlag)
}

func main() {
	flag.Parse()

	cfg, err := getConfig()

	if err != nil {
		log.Fatalln("failed to get config:", err)
		return
	}

	ctx := context.Background()

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

	application := app.New(app.Config{
		Repo:       repo,
		Authorizer: auth,
		Bus:        bus,
	})

	mux := web.NewRouter(application)
	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Addr, cfg.HTTP.Port)

	log.Printf("starting http server (%s)...\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalln("failed to start the server")
	}
}
