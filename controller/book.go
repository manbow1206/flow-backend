package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/model"
	"backend/service"
	"strconv"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad req")
		return
	}
	bookService := service.BookService{}
	err := bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}