package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CoursesRegister(router *gin.RouterGroup) {
	router.GET("/", GetCourses)
	router.POST("/", CreateCourse)
	router.GET("/:id", GetCourse)
	router.PUT("/:id", UpdateCourse)
	router.DELETE("/:id", DeleteCourse)
}

func GetCourses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All Courses",
	})
}
func CreateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a Course",
	})
}

func GetCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a course",
	})
}

func UpdateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a course",
	})
}

func DeleteCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a course",
	})
}
