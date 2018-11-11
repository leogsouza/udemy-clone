package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
	"github.com/leogsouza/udemy-clone/server/courses"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&courses.CourseModel{})
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")

	courses.Register(v1.Group("/courses"))

	r.Run()
}
