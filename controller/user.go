package controller

import (
	"net/http"

	"time"

	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/apd_exam/data"
	"github.com/apd_exam/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func GetUsers(c echo.Context) error {
	db := data.NewDB()
	defer db.Close()

	var users model.User
	if err := db.Model(&model.User{}).Preload("Role").Find(&users).Error; err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Error on retriving users"))

	}

	return c.JSON(http.StatusOK, users)
}

func RegisterAccount(c echo.Context) error {
	var request model.UserRegisterRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	db := data.NewDB()
	defer db.Close()
	var user model.User

	db.Where("email = ?", request.Email).First(&user)

	if user.ID != 0 {
		return errorResponse(c, http.StatusConflict, errors.Wrap(errors.New(""), "The email already exists"))
	}

	userRole, err := GetUserRole(db)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Error on retrive user role"))
	}

	encryptPassword := EncryptPassword(request.Password)
	user.Name = request.Name
	user.Email = request.Email
	user.Password = encryptPassword
	user.Role = []model.Role{userRole}

	if err := db.Create(&user).Error; err != nil {

		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Something wrong when insert user"))
	}

	return c.JSON(http.StatusOK, user)

}

func Login(c echo.Context) error {
	var request model.Login

	if err := c.Bind(&request); err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.Wrap(err, ""))
	}
	db := data.NewDB()
	defer db.Close()
	var user model.User

	request.Password = EncryptPassword(request.Password)

	db.Where("email = ? AND password = ?", request.Email, request.Password).First(&user)

	if user.ID == 0 {
		return errorResponse(c, http.StatusBadRequest, nil)
	}

	accessToken := model.AccessToken{
		Email:     user.Email,
		Token:     randToken(),
		UpdatedAt: time.Now(),
	}

	err := storeToken(db, accessToken)

	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "Error on store Token"))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"email":        accessToken.Email,
		"access_token": accessToken.Token,
	})

}

func storeToken(db *gorm.DB, accessToken model.AccessToken) error {
	var existToken model.AccessToken
	db.Where("email = ?", accessToken.Email).First(&existToken)

	var err error
	if existToken.ID != 0 {
		existToken.Token = accessToken.Token
		existToken.AccessCount += 1
		err = db.Save(&existToken).Error
	} else {
		err = db.Create(&accessToken).Error

	}

	return err

}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%x", b))
	return fmt.Sprintf("%x", h.Sum(nil))
}
