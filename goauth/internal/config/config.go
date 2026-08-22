package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	JWTSecret   string
}

func load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
	if cfg.MongoURI == "" || cfg.MongoDBName == "" || cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("Missing Credentials ! ")
	}
	return cfg, nil
}
