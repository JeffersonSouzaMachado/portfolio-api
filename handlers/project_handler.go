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
		response = ProjectsPortuguese()
	} else {
		response = ProjectsEnglish()
	}

	context.JSON(http.StatusOK, response)
}

func ProjectsPortuguese() []projects.ProjectResponse {
	return []projects.ProjectResponse{
		projects.HelpneiPortuguese(),
	}
}

func ProjectsEnglish() []projects.ProjectResponse {
	return []projects.ProjectResponse{
		projects.HelpneiEnglish(),
	}
}
