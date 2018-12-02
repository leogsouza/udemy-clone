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

// GetAll retrieves all courses in database
func GetAll() ([]CourseModel, error) {
	db := common.GetDB()
	var models []CourseModel

	err := db.Find(&models).Error

	return models, err
}

// GetOne retrieves one course in database
func GetOne(condition interface{}) (CourseModel, error) {
	db := common.GetDB()
	var model CourseModel
	err := db.First(&model).Error

	return model, err
}

// SaveOne saves a course in database
func SaveOne(model *CourseModel) error {
	db := common.GetDB()
	err := db.Save(model).Error

	return err
}

// Update updates a course in database
func (c *CourseModel) Update() error {
	db := common.GetDB()
	err := db.Save(c).Error

	return err
}

// Delete deletes a course in database
func (c *CourseModel) Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(c).Error

	return err
}
