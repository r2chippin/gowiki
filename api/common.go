package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
