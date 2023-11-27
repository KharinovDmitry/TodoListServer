package main

import (
	"TodoListServer/internal/app"
	. "TodoListServer/internal/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	cfgPath := os.Getenv("CONFIG_PATH")
	cfg := MustLoadConfig(cfgPath)

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	application.MustRun()
}

func MustLoadConfig(configPath string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
