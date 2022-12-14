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
	JWT    JWTConfig
	DB     DBConfig
	AMQP   AMQPConfig
}

type ServerConfig struct {
	URL     string
	TimeOut int
}
type JWTConfig struct {
	SecretKey     string
	ExpireMinutes int
}

type DBConfig struct {
	URL         string
	Max         int
	MaxIdle     int
	MaxLifeTime int
}

type AMQPConfig struct {
	URL string
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

	expireMinutes, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	cfg.JWT = JWTConfig{
		SecretKey:     os.Getenv("JWT_SECRET_KEY"),
		ExpireMinutes: expireMinutes,
	}

	cfg.DB = DBConfig{
		URL:         os.Getenv("DB_SERVER_URL"),
		Max:         maxConnection,
		MaxIdle:     maxIdle,
		MaxLifeTime: maxLifeTime,
	}

	cfg.AMQP = AMQPConfig{
		URL: os.Getenv("AMQP_SERVER_URL"),
	}
}

func GetServer() ServerConfig {
	return cfg.Server
}

func GetJWT() JWTConfig {
	return cfg.JWT
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAMQP() AMQPConfig {
	return cfg.AMQP
}
