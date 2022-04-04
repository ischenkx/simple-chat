package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
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
