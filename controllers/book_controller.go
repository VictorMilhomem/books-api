package controllers

import (
	"bookApi/database"
	"bookApi/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	db := database.GetDataBase()
	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not find book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func CreatBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ShowAllBooks(c *gin.Context) {
	db := database.GetDataBase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not list books: " + err.Error(),
		})

		return
	}
	c.JSON(200, books)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDataBase()
	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditBook(c *gin.Context) {
	db := database.GetDataBase()
	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not bind Json " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "could not create the book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

