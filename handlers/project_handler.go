package handlers

import (
	"net/http"

	"portfolio-api/services"

	"github.com/gin-gonic/gin"
)

func ProjectsHandler(context *gin.Context) {
	language := context.DefaultQuery("lang", "en")

	response := services.GetProjects(language)

	context.JSON(http.StatusOK, response)
}
