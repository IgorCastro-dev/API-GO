package main

import (
	"apirest/handlers"
	"apirest/models"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	models.ConnectDatabase()
	e.GET("/", handlers.PaginaPrincipal)
	e.GET("/books", handlers.FindBooks)
	e.GET("/books/:id", handlers.FindBook)
	e.POST("/books", handlers.CreateBook)
	e.PUT("/books/:id", handlers.UpdateBook)
	e.DELETE("/books/:id", handlers.DeleteBook)
	e.Logger.Fatal(e.Start(":8080"))
}
