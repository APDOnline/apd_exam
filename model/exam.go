package model

import "time"

const (
	ExamStatusPending string = "PENDING"
	ExamStatusDone    string = "COMPLETED"
)

type Exam struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Book          Book       `json:"book"`
	BookID        int64      `json:"-"`
	Questions     []Question `json:"questions" gorm:"many2many:exam_question;"`
	QuestionCount int64      `json:"question_count" gorm:"-"`
	Tag           []Tag      `json:"tags" gorm:"many2many:exam_tag;"`
	Status        string     `json:"status"`
	Enable        bool       `json:"enable"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type ExamQuestion struct {
	ExamID     int64 `gorm:"primary_key:true"`
	QuestionID int64 `gorm:"primary_key:true"`
}

type ExamTag struct {
	ExamID int64 `gorm:"primary_key:true"`
	TagID  int64 `gorm:"primary_key:true"`
}

type Tag struct {
	ID   int64  `json:"-" gorm:"primary_key:true"`
	Name string `json:"name"  gorm:"unique"`
}

type CreateExamRequest struct {
	Name   string   `json:"name" form:"name" binding:"required"`
	BookID int64    `json:"book_id" form:"book_id" binding:"required"`
	Tag    []string `json:"tag" form:"tag"`
}

type UpdateQuestionToExamRequest struct {
	ExamID     int64 `json:"exam_id" binding:"required"`
	QuestionID int64 `json:"question_id" binding:"required"`
}
