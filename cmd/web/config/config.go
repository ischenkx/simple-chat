package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Postgres struct {
		URL string `json:"url" yaml:"url"`
	} `json:"postgres" yaml:"postgres"`

	JWT struct {
		Key            string `json:"key" yaml:"key"`
		ExpirationTime int64  `json:"expiration_time" yaml:"expiration_time"`
	} `json:"jwt" yaml:"jwt"`

	HTTP struct {
		Addr string `json:"addr" yaml:"addr"`
		Port uint16 `json:"port" yaml:"port"`
	}
}

func FromFile(filename string) (Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}

	if strings.HasSuffix(filename, ".json") {
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return config, err
		}
	} else {
		if err := yaml.NewDecoder(file).Decode(&config); err != nil {
			return config, err
		}
	}
	return config, nil
}

func FromENV() (Config, error) {
	var config Config

	// DB
	config.Postgres.URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_ADDR"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	// Server
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return config, err
	}

	config.HTTP.Port = uint16(port)
	config.HTTP.Addr = os.Getenv("ADDR")

	// JWT
	config.JWT.Key = os.Getenv("JWT_KEY")
	expTime, err := time.ParseDuration(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		return config, err
	}
	config.JWT.ExpirationTime = expTime.Milliseconds()

	return config, nil
}
