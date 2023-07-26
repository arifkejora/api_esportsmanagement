package Attendance

import (
	"api_esportsmanagement/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler untuk menandai kehadiran oleh peserta
func MarkAttendance(c *gin.Context) {
	var attendance model.MAttendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: Simpan data kehadiran ke database
	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked successfully"})
}
