package courses

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	courses, err := GetAllCourses()

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

	c.JSON(200, course)
}

// GetCourse retrieve a course
func GetCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a course",
	})
}

// UpdateCourse updates a course
func UpdateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a course",
	})
}

// DeleteCourse deletes a course
func DeleteCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a course",
	})
}
