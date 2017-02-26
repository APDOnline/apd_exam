package data

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/apd_exam/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type DatabaseStruct struct {
	Name     string
	User     string
	Password string
	IP       string
}

func (d *DatabaseStruct) ToConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", d.IP, d.User, d.Name, d.Password)
}

var DatabaseInfo DatabaseStruct

func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", DatabaseInfo.ToConnectionString())

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.SingularTable(true)
	return db

}

func CreateSchema(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Role{}, &model.User{}, &model.AccessToken{}).Error; err != nil {
		return errors.Wrap(err, "Migrate schema failes")
	}

	if err := db.AutoMigrate(&model.Book{}, &model.Question{}, &model.Option{}, &model.Reference{}, &model.Difficulty{}).Error; err != nil {
		return errors.Wrap(err, "Migrate schema failes")
	}

	if err := db.AutoMigrate(&model.Exam{}, &model.ExamTag{}, &model.ExamQuestion{}, &model.ExamTag{}, &model.Tag{}).Error; err != nil {
		return errors.Wrap(err, "Migrate schema failes")
	}

	log.Debug("Create database schema completed")

	return nil
}

func DefaultData(db *gorm.DB) error {

	var adminRole model.Role
	var userRole model.Role
	db.Where(model.Role{
		Authority: "ROLE_ADMIN",
	}).FirstOrCreate(&adminRole)

	db.Where(model.Role{
		Authority: "ROLE_USER",
	}).FirstOrCreate(&userRole)

	adminUser := &model.User{
		Name:     "admin",
		Email:    "admin",
		Password: "admin",
		Role: []model.Role{
			adminRole,
		},
	}

	err := db.Where("email = ?", adminUser.Email).FirstOrCreate(&adminUser).Error

	if err != nil {
		return errors.Wrap(err, "Retrive Admin user failed")
	}

	user1 := &model.User{
		Name:     "Jay",
		Email:    "jack08300@gmail.com",
		Password: "aaaaaa",
		Role:     []model.Role{userRole},
	}
	err = db.Where("email = ?", user1.Email).FirstOrCreate(&user1).Error

	if err != nil {
		return errors.Wrap(err, "Retrive Admin user failed")
	}

	//Create Difficulty
	difficulty := []model.Difficulty{
		{
			Name:  "EASY",
			Point: 0.3,
		},
		{
			Name:  "MEDIUM",
			Point: 0.5,
		},
		{
			Name:  "HARD",
			Point: 0.8,
		},
	}

	for _, d := range difficulty {
		if err := db.Where("name = ?", d.Name).FirstOrCreate(&d).Error; err != nil {
			return errors.Wrap(err, "Error on creating difficulty")
		}
	}

	log.Debug("Default users creation completed")

	return nil
}
