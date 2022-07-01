package handlers

import (
	"apirest/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /
func PaginaPrincipal(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func FindBooks(c echo.Context) error {
	var books []models.Book
	if err := models.DB.Find(&books).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string][]models.Book{"data": books})
}

//GET /books
func FindBook(c echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]models.Book{"data": book})
}

//POST /books
func CreateBook(c echo.Context) error {
	var input models.CreateBook
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}
	var book = models.Book{Title: input.Title, Author: input.Author}
	if err := models.DB.Create(&book).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"Erro": "erro no servidor"})
	}
	return c.JSON(http.StatusOK, map[string]models.Book{"data": book})
}

func UpdateBook(c echo.Context) error {
	var input models.UpdateBook
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}

	if err := models.DB.Model(&book).Updates(input).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"Erro": "erro no servidor"})
	}

	return c.JSON(http.StatusOK, map[string]models.Book{"data": book})
}

func DeleteBook(c echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Erro": err.Error()})
	}
	if err := models.DB.Delete(&book).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"Erro": "erro no servidor"})
	}
	return c.JSON(http.StatusOK, map[string]bool{"Data": true})
}
