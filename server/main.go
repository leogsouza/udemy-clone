package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/udemy-clone/server/common"
	"github.com/leogsouza/udemy-clone/server/courses"
	"github.com/leogsouza/udemy-clone/server/users"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&courses.CourseModel{})
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")

	courses.Register(v1.Group("/courses"))
	users.Register(v1.Group("/users"))

	r.Run()
}
