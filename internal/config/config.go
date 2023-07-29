package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func GetConfig() *Config {
	conf := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
	if conf.DBUser == "" || conf.DBPassword == "" || conf.DBHost == "" || conf.DBPort == "" || conf.DBName == "" {
		log.Fatalf("Environment variables for DB not set properly")
	}
	
	return conf
}
