package main

import (
	"api/players"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("POSTSQL_URL")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	r := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	fport := fmt.Sprintf("http://localhost:%s", os.Getenv("FRONT_PORT"))
	config.AllowOrigins = []string{fport} // This needs to be Frontend URL
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.POST("/players", func(c *gin.Context) {
		var newPlayer players.Player

		if err := c.BindJSON(&newPlayer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		hashedPassword, err := players.HashPassword(newPlayer.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		newPlayer.Password = hashedPassword

		result := db.Create(&newPlayer)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
			return
		}

		c.JSON(http.StatusCreated, newPlayer)
	})

	var fplayers []players.APIPlayer

	db.Model(&players.Player{}).Find(&fplayers)

	r.GET("/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": fplayers,
		})
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}

/*
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("POSTSQL_URL")

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	//newplayer := players.CreatePlayer()

	//db.AutoMigrate(&players.Player{})
	//db.Create(&players.Player{Username: newplayer.Username, Password: newplayer.Password, CharName: newplayer.CharName})

	var fplayers []players.APIPlayer

	db.Model(&players.Player{}).Find(&fplayers)

	r := gin.Default()

	r.POST("/players", func(c *gin.Context) {
		var newPlayer players.Player

		if err := c.BindJSON(&newPlayer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Hash the password before storing in the database
		hashedPassword, err := players.HashPassword(newPlayer.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		newPlayer.Password = hashedPassword

		// Create the player record in the database
		result := db.Create(&newPlayer)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
			return
		}

		c.JSON(http.StatusCreated, newPlayer)
	})

	r.GET("/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": fplayers,
		})
	})
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}

*/
