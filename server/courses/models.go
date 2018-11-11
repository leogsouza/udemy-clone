package courses

import (
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
)

// CourseModel hols a course struct
type CourseModel struct {
	gorm.Model
	Title       string
	Description string
	Cover       string
}

// GetAllCourses retrieves all courses in database
func GetAllCourses() ([]CourseModel, error) {
	db := common.GetDB()
	var models []CourseModel

	err := db.Find(&models).Error

	return models, err
}

func SaveOne(c *CourseModel) error {
	db := common.GetDB()
	err := db.Save(c).Error

	return err
}
