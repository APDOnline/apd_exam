package model

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string `gorm:"unique"`
	Password  string `json:"-"`
	Role      []Role `gorm:"many2many:user_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Role struct {
	ID        int64  `json:"-" gorm:"AUTO_INCREMENT"`
	Authority string `json:"authority" gorm:"unique"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessToken struct {
	ID          int64     `gorm:"AUTO_INCREMENT"`
	Email       string    `gorm:"unique;not null"`
	Token       string    `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	AccessCount int
}
