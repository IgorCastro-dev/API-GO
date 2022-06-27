package main

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"dado": "Ola mundo"})
	})
	models.ConnectDatabase()
	r.Run()
}
