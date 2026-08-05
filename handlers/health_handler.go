package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealtHandler(context *gin.Context) {
	response := HealthResponse{
		Status: "ok",
	}

	context.JSON(http.StatusOK, response)
}
