package config

import (
	"log"
	"reflect"

	dotenv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// @TODO implement a good configuration service.
type (
	// config wraps all configuration parts.
	config struct {
		App *AppConfig
		Db  *DbConfig
	}

	// AppConfig describes application config.
	AppConfig struct {
		Host string
		Port int
	}

	// DbConfig describes database config.
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
	// configuration keeps config instance.
	configuration = &config{
		App: &AppConfig{},
		Db:  &DbConfig{},
	}
)

// App returns application config.
func App() AppConfig {
	return *configuration.App
}

// Db returns database config.
func Db() DbConfig {
	return *configuration.Db
}

// Load loads and inits the app configuration.
func Load() error {
	log.Println("Load configuration...")

	// load options from .env file to environment
	if err := dotenv.Load(); err != nil {
		return err
	}

	// init config using environment
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
