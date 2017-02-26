package model

import "time"

type Book struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions,omitempty" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Question struct {
	ID              int64      `json:"id" gorm:"AUTO_INCREMENT"`
	Type            string     `json:"type"`
	Text            string     `json:"text"`
	Image           string     `json:"image"`
	Category        string     `json:"category"`
	Difficulty      Difficulty `json:"difficulty,omitempty"`
	DifficultyID    int64      `json:"-"`
	LearningOutcome string     `json:"learning_outcome"`
	LONumber        string     `json:"lo_number"`
	Options         []Option   `json:"options,omitempty"`
	Reference       Reference  `json:"reference,omitempty"`
	AnswerExp       string     `json:"answer_exp"`
	ChoiceLen       int64      `json:"choice_len"`
	QuestionLen     int64      `json:"question_len"`
	ReadLoad        int32      `json:"read_load"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	BookID          int64
}

type Option struct {
	ID         int64   `json:"id"`
	Text       string  `json:"text"`
	Type       string  `json:"type"`
	Image      string  `json:"image"`
	Point      float32 `json:"point" gorm:"type:numeric(3,2)"`
	QuestionID int64   `json:"-"`
}

type Reference struct {
	ID           int64  `json:"-"`
	ChapterID    int    `json:"chapter_id"`
	ChapterTitle string `json:"chapter_title"`
	PageRef      string `json:"page_ref"`
	QuestionID   int64  `json:"-"`
}

type Difficulty struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Point float32 `json:"point" gorm:"type:numeric(3,2)"`
}
