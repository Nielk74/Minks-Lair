package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func mergeEnv() map[string]string {
	mergedConfig := make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			mergedConfig[pair[0]] = pair[1]
		}
	}
	envFile, _ := godotenv.Read(".env")
	for key, value := range envFile {
		if value != "" {
			mergedConfig[key] = value
		} else {
			mergedConfig[key] = os.Getenv(key)
		}
	}
	return mergedConfig
}

func (c *Config) testConnection() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database!")
}

func RunTestConnection() {
	envConfig := mergeEnv()
	log.Println("Connecting to database...")
	config := Config{
		Host:     envConfig["DB_HOST"],
		Port:     envConfig["DB_PORT"],
		User:     envConfig["DB_USER"],
		Password: envConfig["DB_PASSWORD"],
		Database: envConfig["DB_NAME"],
	}
	config.testConnection()
}

func main() {
	RunTestConnection()
}
