package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string
	DBAddr   string
}

func NewConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("fib", &conf)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &conf, nil
}
