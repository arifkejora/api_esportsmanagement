package main

import (
	"api_esportsmanagement/database"
	"api_esportsmanagement/router"
	"log"
)

func main() {
	database.ConnectDB()
	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
