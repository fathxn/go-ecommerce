package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	Name      string
	DBAddress string
}

var Envs = initConfig()

func initConfig() DBConfig {
	godotenv.Load()

	return DBConfig{
		Host:      getEnv("DB_HOST", "http://localhost"),
		Port:      getEnv("DB_PORT", "8080"),
		User:      getEnv("DB_USER", "root"),
		Password:  getEnv("DB_PASSWORD", ""),
		Name:      getEnv("DB_NAME", "go_ecommerce"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
