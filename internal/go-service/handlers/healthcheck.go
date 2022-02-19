package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthcheckHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
