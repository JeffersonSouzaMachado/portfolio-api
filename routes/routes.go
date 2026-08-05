package routes

import (
	"portfolio-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealtHandler)

	router.GET("/metrics", handlers.MetricsHandler)

	router.GET("/projects", handlers.ProjectsHandler)

	router.GET("/project/:id", handlers.HealtHandler)

	router.GET("/skills", handlers.HealtHandler)

	router.GET("/contact", handlers.HealtHandler)

	router.POST("/send-message", handlers.HealtHandler)
}
