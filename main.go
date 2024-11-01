package main

import (
	"fmt"
	"log"
	"net/http"
	"todoist/db"
	"todoist/router"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up the router
	r := router.SetupRouter()

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
