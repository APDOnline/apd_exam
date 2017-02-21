package controller

import (
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/apd_exam/model"
	"github.com/jinzhu/gorm"
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
