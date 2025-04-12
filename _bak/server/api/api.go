package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api/controller"
	"strings"
)

func BuildRouter(spa http.FileSystem) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Authorization", "Origin"}
	router.Use(cors.New(corsConfig))
	api := router.Group("/api")
	{
		api.GET("/telemetry", controller.GetTelemetries)
		api.POST("/telemetry", controller.CreateTelemetry)

		api.GET("/experience", controller.GetExperiences)
		api.POST("/experience", controller.CreateExperience)
		api.DELETE("/experience/:id", controller.DeleteExperience)
		api.PUT("/experience", controller.UpdateExperience)

		api.GET("/technology", controller.GetTechnologies)
		api.POST("/technology", controller.CreateTechnology)
		api.DELETE("/technology/:id", controller.DeleteTechnology)
		api.PUT("/technology", controller.UpdateTechnology)

		api.POST("/authentication", controller.AuthenticateUser)
	}
	if spa != nil {
		// all routes other than api must be spa, get requested file from http.FS
		router.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			// in case a non-existing api route is hit, thus falling here, abort with 404
			if strings.HasPrefix(path, "/api") {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			c.FileFromFS(path, spa)
		})
	}

	return router
}
