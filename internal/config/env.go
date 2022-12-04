package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var cfg *config

type config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	URL     string
	TimeOut int
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

	readTimeOut, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	maxConnection, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	cfg = new(config)
	cfg.Server = ServerConfig{
		URL:     os.Getenv("SERVER_URL"),
		TimeOut: readTimeOut,
	}

	cfg.DB = DBConfig{
		URL:         os.Getenv("DB_SERVER_URL"),
		Max:         maxConnection,
		MaxIdle:     maxIdle,
		MaxLifeTime: maxLifeTime,
	}
}

func GetServer() ServerConfig {
	return cfg.Server
}

func GetDB() DBConfig {
	return cfg.DB
}
