package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var cfg *config

type config struct {
	DB DBConfig
}

type DBConfig struct {
	URL         string
	Max         int
	MaxIdle     int
	MaxLifeTime int
}

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	maxConnection, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	cfg = new(config)
	cfg.DB = DBConfig{
		URL:         os.Getenv("DB_SERVER_URL"),
		Max:         maxConnection,
		MaxIdle:     maxIdle,
		MaxLifeTime: maxLifeTime,
	}
}

func GetDB() DBConfig {
	return cfg.DB
}
