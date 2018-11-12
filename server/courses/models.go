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
func GetAll() ([]CourseModel, error) {
	db := common.GetDB()
	var models []CourseModel

	err := db.Find(&models).Error

	return models, err
}

func GetOne(condition interface{}) (CourseModel, error) {
	db := common.GetDB()
	var model CourseModel
	err := db.First(&model).Error

	return model, err
}

func SaveOne(model *CourseModel) error {
	db := common.GetDB()
	err := db.Save(model).Error

	return err
}

func (c *CourseModel) Update() error {
	db := common.GetDB()
	err := db.Save(c).Error

	return err
}

func (c *CourseModel) Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(c).Error

	return err
}
