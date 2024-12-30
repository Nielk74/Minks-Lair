package database

import (
	"cosmink/libs/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetConnection() *sql.DB {
	envConfig := utils.GetEnv()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envConfig["DB_HOST"], envConfig["DB_PORT"], envConfig["DB_USER"], envConfig["DB_PASSWORD"], envConfig["DB_NAME"])
	log.Println(fmt.Sprintf("Connecting to database: %s", connStr))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (c *Config) testConnection() bool {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return false
	}

	log.Println("Successfully connected to database!")
	return true
}

func RunTestConnection() {
	envConfig := utils.GetEnv()
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
