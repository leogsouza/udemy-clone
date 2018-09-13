package models

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Title       string
	Description string
	Cover       string
}

func GetAllCourses() []Course {

}
