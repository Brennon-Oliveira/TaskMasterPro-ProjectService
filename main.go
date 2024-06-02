package main

import (
	"Brennon-Oliveira/TaskMasterPro-ProjectService/internal/handlers"
	"Brennon-Oliveira/TaskMasterPro-ProjectService/pkg/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	dbConn := db.InitPostgres()

	handlers.RegistryRoutes(r, dbConn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}