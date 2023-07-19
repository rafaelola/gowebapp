package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var database *sql.DB

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func Connect() *sql.DB {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	connectionDetails := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)
	database, err := sql.Open("mysql", connectionDetails)
	if err != nil {
		log.Println("Could not connect!")
		panic(err)
	}

	log.Println("Connected.")
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	return database

}
