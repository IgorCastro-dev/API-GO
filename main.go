package main

import (
	"api/handlers"
	"api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/", handlers.PaginaPrincipal)
	r.GET("/books", handlers.FindBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.FindBook)
	r.Run()
}
