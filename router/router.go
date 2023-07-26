package router

import (
	"api_esportsmanagement/controller.go/Activity"
	"api_esportsmanagement/controller.go/Attendance"
	"api_esportsmanagement/controller.go/Event"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes for admin and participant
	admin := r.Group("/admin")
	{
		admin.POST("/create_event", Event.CreateEvent)
		// Add other admin routes here
	}

	participant := r.Group("/participant")
	{
		participant.POST("/attendance", Attendance.MarkAttendance)
		participant.POST("/upload_activity", Activity.UploadActivity)
		// Add other participant routes here
	}

	return r
}
