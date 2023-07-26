package Event

import (
	"api_esportsmanagement/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler untuk membuat acara oleh admin
func CreateEvent(c *gin.Context) {
	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: Simpan data acara ke database menggunakan model
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}
