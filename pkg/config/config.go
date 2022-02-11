package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string `default:"8080"`
	DBAddr   string `default:"localhost:6379"`
}

func NewConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("fibonaccisrv", &conf)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &conf, nil
}
