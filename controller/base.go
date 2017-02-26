package controller

import (
	"crypto/sha256"
	"fmt"
	"io"

	"time"

	"github.com/Sirupsen/logrus"
	"github.com/apd_exam/model"
	"github.com/jinzhu/gorm"
	"github.com/kidsdynamic/childrenlab_v2/constants"
	"github.com/labstack/echo"
)

type ErrResp struct {
	Message string
	Error   error
}

func errorResponse(c echo.Context, httpStatus int, err error) error {
	var errResp ErrResp
	if err != nil {
		logrus.Errorf("%+v", err)
		errResp = ErrResp{
			Message: err.Error(),
		}
	}

	return c.JSON(httpStatus, errResp)

}

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)

	return fmt.Sprintf("%x", h.Sum(nil))

}

func GetUserRole(db *gorm.DB) (model.Role, error) {
	var role model.Role
	if err := db.Where("authority = ?", "ROLE_USER").First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func GetNowTime() time.Time {
	now := time.Now()

	s := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	t, err := time.Parse(constants.TimeLayout, s)

	if err != nil {
		fmt.Printf("Error on get now time. %#v", err)

	}

	return t
}
