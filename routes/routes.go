package routes

import (
	"os"

	"github.com/leozz37/glucose-db-save/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.GET("/save", handlers.SaveGlucose)
	router.NoRoute(noRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"status": 404, "message": "Page not found"})
}
