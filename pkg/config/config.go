package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string
	DBAddr   string
	User     string
}

func NewConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("fib", &conf)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	fmt.Println(conf.HTTPAddr, conf.DBAddr, conf.User)

	return &conf, nil
}
