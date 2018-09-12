package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leogsouza/udemy-clone/server/courses"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api")

	courses.CoursesRegister(v1.Group("/courses"))

	r.Run()
}
