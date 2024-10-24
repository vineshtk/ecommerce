package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	var err error
	connStr := "postgres://ecommerce_user:password@localhost:5432/ecommerce?sslmode=disable"
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}
	fmt.Println("Database connected successfully!")
}
