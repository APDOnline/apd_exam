package controller

import (
	"net/http"

	"fmt"

	"github.com/apd_exam/data"
	"github.com/apd_exam/model"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func GetAllBooks(c echo.Context) error {

	db := data.NewDB()
	defer db.Close()

	books := &[]model.Book{}

	if err := db.Find(&books).Order("created_at desc").Error; err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Error on retrive books"))
	}

	return c.JSON(http.StatusOK, books)
}

func GetQuestionsByBook(c echo.Context) error {
	bookID := c.Param("bookID")

	db := data.NewDB()
	defer db.Close()

	questions := &[]model.Question{}
	if err := db.Model(&model.Question{}).Preload("Difficulty").Preload("Reference").Preload("Options").Where("book_id = ?", bookID).Limit(10).Find(&questions).Error; err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf("Error on retrive questions by book ID: %s", bookID)))
	}

	return c.JSON(http.StatusOK, questions)
}
