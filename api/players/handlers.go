package players

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreatePlayer(c *gin.Context) {
	var newPlayer Player

	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hashedPassword, err := HashPassword(newPlayer.Password)
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
}

func GetPlayers(c *gin.Context) {
	var fplayers []Player
	db.Model(&Player{}).Find(&fplayers)
	c.JSON(http.StatusOK, gin.H{"data": fplayers})
}
