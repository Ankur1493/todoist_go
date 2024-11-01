package db

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
    connStr := "user=postgres password=postgres dbname=todoist host=localhost port=5432 sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Cannot connect to the database:", err)
    }
    log.Println("Connected to the database")
}

