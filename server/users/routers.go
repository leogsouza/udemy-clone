package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
)

// Register contains all router for user api
func Register(router *gin.RouterGroup) {
	router.GET("/", GetUsers)
	router.POST("/", CreateUser)
	router.GET("/:id", GetUser)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DeleteUser)
}

// GetUsers outputs all users
func GetUsers(c *gin.Context) {
	users, err := GetAll()

	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	if users == nil {
		users = make([]UserModel, 0)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

// CreateUser creates an user
func CreateUser(c *gin.Context) {
	var user UserModel
	c.BindJSON(&user)

	err := SaveOne(&user)
	if err != nil {
		log.Fatalln("failed to save the user:", err)
	}

	c.JSON(http.StatusOK, user)
}

// GetUser retrieve an user
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", err))
		return
	}

	user, err := GetOne(&UserModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", err))
		return
	}
	user, err := GetOne(&UserModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", err))
	}
	c.BindJSON(&user)
	user.Update()

	c.JSON(http.StatusOK, user)
}

// DeleteUser removes an user
func DeleteUser(c *gin.Context) {
	var user UserModel

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", err))
		return
	}

	err = user.Delete(&UserModel{Model: gorm.Model{ID: uint(id)}})
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete an user",
	})
}
