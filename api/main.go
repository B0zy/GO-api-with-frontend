package main

import (
	"api/database"
	"api/players"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	db := database.GetDB()
	players.SetDB(db) // Set the database connection in the players package

	r := gin.Default()

	config := cors.DefaultConfig()
	fport := fmt.Sprintf("http://localhost:%s", os.Getenv("FRONT_PORT"))
	config.AllowOrigins = []string{fport}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.POST("/players", players.CreatePlayer)
	r.GET("/players", players.GetPlayers)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
