package courses

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
)

// Register contains all routes for courses
func Register(router *gin.RouterGroup) {
	router.GET("/", GetCourses)
	router.POST("/", CreateCourse)
	router.GET("/:id", GetCourse)
	router.PUT("/:id", UpdateCourse)
	router.DELETE("/:id", DeleteCourse)
}

// GetCourses outputs all courses
func GetCourses(c *gin.Context) {
	courses, err := GetAll()

	if err != nil {
		log.Fatalln("failed to connect to mongo:", err)
	}

	if courses == nil {
		courses = make([]CourseModel, 0)
	}
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

// CreateCourse creates a course
func CreateCourse(c *gin.Context) {
	var course CourseModel
	c.BindJSON(&course)

	err := SaveOne(&course)
	if err != nil {
		log.Fatalln("failed to connect to mongo:", err)
	}

	c.JSON(http.StatusOK, course)
}

// GetCourse retrieve a course
func GetCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("courses", err))
		return
	}
	course, err := GetOne(&CourseModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("courses", err))
		return
	}

	c.JSON(http.StatusOK, course)
}

// UpdateCourse updates a course
func UpdateCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("courses", err))
		return
	}
	course, err := GetOne(&CourseModel{Model: gorm.Model{ID: uint(id)}})
	c.BindJSON(&course)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("courses", err))
		return
	}
	course.Update()

	c.JSON(http.StatusOK, course)
}

// DeleteCourse deletes a course
func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var course CourseModel
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("courses", err))
		return
	}
	err = course.Delete(&CourseModel{Model: gorm.Model{ID: uint(id)}})
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a course",
	})
}
