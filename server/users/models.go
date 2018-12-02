package users

import (
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
)

// UserModel represents a User entity in database
type UserModel struct {
	gorm.Model
	Name     string
	Username string
	Email string
	Password string
}

// GetAll retrieves all users in database
func GetAll() ([]UserModel, error) {
	db := common.GetDB()
	var models []UserModel

	err := db.Find(&models).Error

	return models, err
}

// GetOne retrieves one user in database
func GetOne(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.First(&model).Error

	return model, err
}

// SaveOne saves a user in database
func SaveOne(model *UserModel) error {
	db := common.GetDB()
	err := db.Save(model).Error

	return err
}

// Update udpates a user in database
func (u *UserModel) Update() error {
	db := common.GetDB()
	err := db.Save(u).Error
	return err
}

// Delete removes a user in database
func (u *UserModel) Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(u).Error

	return err
}
