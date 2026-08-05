package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetricsResponse struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func MetricsHandler(context *gin.Context) {
	language := context.DefaultQuery("lang", "en")

	var response []MetricsResponse

	if language == "pt" {
		response = metricsPortuguese()
	} else {
		response = metricsEnglish()
	}

	context.JSON(http.StatusOK, response)
}

func metricsPortuguese() []MetricsResponse {
	return []MetricsResponse{
		{
			Title: "04+",
			Text:  "anos de experiência",
		},
		{
			Title: "14",
			Text:  "Apps em produção",
		},
		{
			Title: "5.6M",
			Text:  "usuários ativos",
		},
		{
			Title: "110+",
			Text:  "PRs Mergeados/ME",
		},
	}
}

func metricsEnglish() []MetricsResponse {
	return []MetricsResponse{
		{
			Title: "04+",
			Text:  "years experience",
		},
		{
			Title: "14",
			Text:  "production apps",
		},
		{
			Title: "5.6M",
			Text:  "active users",
		},
		{
			Title: "110+",
			Text:  "prs merged/mo",
		},
	}
}
