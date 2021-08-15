package handlers

import (
	"github.com/leozz37/glucose-db-save/db"

	"github.com/gin-gonic/gin"
)

func SaveGlucose(c *gin.Context) {
	glucose := c.Query("last")
	gap := c.Query("gap")

	db.Save(glucose, gap)

	c.JSON(200, gin.H{"message": "success"})
}
