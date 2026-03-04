package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
    MongoURI     string
    MongoDbName  string
    ServerPort   string
}

func getEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return val, nil
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %v", err)
	}
	mongoUri, err := getEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}
	port, err := getEnv("PORT")
	if err != nil {
		return Config{}, err
	}
	MongoDbName, err := getEnv("MONGO_DB_NAME")
	fmt.Printf("Loaded Config - MongoURI: %s, Port: %s, MongoDbName: %s\n", mongoUri, port, MongoDbName)
	if err != nil {
		return Config{}, err
	}
	return Config{
		MongoURI:   mongoUri,
		MongoDbName:  MongoDbName,
		ServerPort:   port,
	}, nil
}
