package handlers

import (
	"net/http"

	"portfolio-api/projects"

	"github.com/gin-gonic/gin"
)

func ProjectsHandler(context *gin.Context) {
	language := context.DefaultQuery("lang", "en")

	var response []projects.ProjectResponse

	if language == "pt" {
		response = projects.ProjectsPortuguese()
	} else {
		response = projects.ProjectsEnglish()
	}

	context.JSON(http.StatusOK, response)
}
