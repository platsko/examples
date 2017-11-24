// Copyright © 2020 The EVEN Lab Team

package config

import (
	"log"
	"reflect"

	dotenv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// @TODO implement a good configuration service.
type (
	config struct {
		APP *AppConfig
		DB  *DbConfig
	}

	AppConfig struct {
		Host string
		Port int
	}

	DbConfig struct {
		Code string
		Conn string
		Host string
		Name string
		Port int
		Pass string
		User string
	}
)

var (
	configuration = &config{
		APP: &AppConfig{},
		DB:  &DbConfig{},
	}
)

// App returns application config.
func APP() AppConfig {
	return *configuration.APP
}

// DB returns database config.
func DB() DbConfig {
	return *configuration.DB
}

// Load loads and inits the app configuration.
func Load() error {
	log.Println("Load configuration...")

	// Load options from .env file to environment.
	if err := dotenv.Load(); err != nil {
		return err
	}

	// Init config using environment.
	if err := initConfig(); err != nil {
		return err
	}

	return nil
}

// initConfig inits config using environment.
func initConfig() error {
	conf := reflect.ValueOf(configuration).Elem()
	l := conf.NumField()
	for idx := 0; idx < l; idx++ {
		part := conf.Field(idx)
		name := conf.Type().Field(idx).Name
		if err := envconfig.Process(name, part.Interface()); err != nil {
			return err
		}
	}

	return nil
}
