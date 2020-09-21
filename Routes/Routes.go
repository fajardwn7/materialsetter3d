package routes

import (
	"materialsetter3d/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("stuffs", controllers.GetStuffs)
	}
	return r
}
