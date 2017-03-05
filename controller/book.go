package controller

import (
	"net/http"

	"fmt"

	"github.com/apd_exam/cache"
	"github.com/apd_exam/data"
	"github.com/apd_exam/model"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const BookCacheKey = "BOOK_LIST"
const BookQuestionKey = "BOOK_QUESTION_LIST"

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
	questions := &[]model.Question{}
	cacheKey := fmt.Sprintf("%s_%d", BookQuestionKey, bookID)

	questionCache, ok := cache.Cache.Get(cacheKey)
	if !ok {
		db := data.NewDB()
		defer db.Close()

		if err := db.Model(&model.Question{}).Preload("Difficulty").Preload("Reference").Preload("Options").Where("book_id = ?", bookID).Find(&questions).Error; err != nil {
			return errorResponse(c, http.StatusInternalServerError, errors.Wrapf(err, "Error on retrive questions by book ID: %s", bookID))
		}

		cache.Cache.Set(cacheKey, questions)
	} else {
		questions = questionCache.(*[]model.Question)
	}

	return c.JSON(http.StatusOK, questions)
}

func CleanBookCache(c echo.Context) error {
	cache.Cache.Remove(BookCacheKey)
	return nil
}
