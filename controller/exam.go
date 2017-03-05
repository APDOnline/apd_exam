package controller

import (
	"net/http"

	"fmt"

	"github.com/apd_exam/data"
	"github.com/apd_exam/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func CreateExam(c echo.Context) error {
	request := &model.CreateExamRequest{}
	if err := c.Bind(&request); err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Error on unmarshall reqeust"))
	}

	db := data.NewDB()
	defer db.Close()

	exam := &model.Exam{}
	db.Where("name = ?", request.Name).First(&exam)

	if exam.ID != 0 {
		return errorResponse(c, http.StatusConflict, errors.Wrap(errors.New("The exam exists in database"), ""))
	}

	tags, err := createTag(db, request.Tag)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Error on creating tags"))
	}

	var book model.Book
	if err := db.Where("id = ?", request.BookID).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "The book is not exist"))
		} else {
			return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Error on getting book"))
		}
	}

	exam = &model.Exam{
		Name:   request.Name,
		Book:   book,
		BookID: book.ID,
		Status: model.ExamStatusPending,
		Tag:    tags,
	}

	if err := db.Create(&exam).Error; err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrapf(err, "Error on creating exam: %#v", exam))
	}

	return c.JSON(http.StatusOK, exam)
}

func UpdateQuestionToExam(c echo.Context) error {
	var request model.UpdateQuestionToExamRequest

	if err := c.Bind(&request); err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Error on parse parameters"))
	}

	db := data.NewDB()
	defer db.Close()

	var exam model.Exam
	if err := db.Where("id = ?", request.ExamID).Preload("Book").Preload("Tag").First(&exam).Error; err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Error on Finding questions"))
	}

	for _, question := range exam.Questions {
		if question.ID == request.QuestionID {
			return errorResponse(c, http.StatusConflict, errors.Wrap(errors.New("Conflict"), "The record already exists"))
		}
	}

	var question model.Question

	if err := db.Where("id = ?", request.QuestionID).First(&question).Error; err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "Error on Finding questions"))
	}

	exam.Questions = append(exam.Questions, question)
	if err := db.Omit("Book", "Tag").Save(&exam).Error; err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrapf(err, "Error on saving exam.\n%#v", exam))
	}

	return c.JSON(http.StatusOK, exam)

}

func GetExam(c echo.Context) error {
	examID := c.Param("id")

	db := data.NewDB()
	defer db.Close()

	var exam model.Exam
	if err := db.Where("id = ?", examID).Preload("Book").Preload("Questions").Preload("Questions.Options").Preload("Questions.Difficulty").Preload("Questions.Reference").Preload("Tag").First(&exam).Error; err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrapf(err, "Error on retrive exam. %#v", exam))
	}

	return c.JSON(http.StatusOK, exam)
}

func getQuestionListByIDs(db *gorm.DB, bookId int64, questionIds []int64) ([]model.Question, error) {
	var questions []model.Question

	if err := db.Where("book_id = ? AND id in (?)", bookId, questionIds).Find(&questions).Error; err != nil {
		return questions, err
	}
	return questions, nil
}

func createTag(db *gorm.DB, stringTags []string) ([]model.Tag, error) {
	var tags []model.Tag
	for _, stringTag := range stringTags {
		var tag model.Tag
		if err := db.FirstOrCreate(&tag, model.Tag{Name: stringTag}).Error; err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}
	fmt.Printf("%#v", tags)

	return tags, nil
}
