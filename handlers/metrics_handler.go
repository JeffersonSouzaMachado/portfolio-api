package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetricsResponse struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Color string `json:"color"`
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
			Color: "seconday",
		},
		{
			Title: "14",
			Text:  "Apps em produção",
			Color: "accent",
		},
		{
			Title: "5.6M",
			Text:  "usuários ativos",
			Color: "secondaryFixedDim",
		},
		{
			Title: "110+",
			Text:  "PRs Mergeados/ME",
			Color: "errorContainer",
		},
	}
}

func metricsEnglish() []MetricsResponse {
	return []MetricsResponse{
		{
			Title: "04+",
			Text:  "years experience",
			Color: "seconday",
		},
		{
			Title: "14",
			Text:  "production apps",
			Color: "accent",
		},
		{
			Title: "5.6M",
			Text:  "active users",
			Color: "secondaryFixedDim",
		},
		{
			Title: "110+",
			Text:  "prs merged/mo",
			Color: "errorContainer",
		},
	}
}
