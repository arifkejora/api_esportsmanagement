package Activity

import (
	"api_esportsmanagement/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler untuk mengunggah kegiatan oleh peserta
func UploadActivity(c *gin.Context) {
	var activity model.MActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// TODO: Simpan data kegiatan ke database menggunakan model
	c.JSON(http.StatusOK, gin.H{"message": "Activity uploaded successfully"})
}
