package common

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../learning.db")
	if err != nil {
		fmt.Println("db err", err)
	}

	db.DB().SetMaxIdleConns(10)

	DB = db
	return DB
}

func TestDBInit() *gorm.DB {
	testDb, err := gorm.Open("sqlite3", "./../learning_test.db")
	if err != nil {
		fmt.Println("db err", err)
	}

	testDb.DB().SetMaxIdleConns(3)
	testDb.LogMode(true)
	DB = testDb
	return DB
}

func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../learning_test.db")
	return err
}

func GetDB() *gorm.DB {
	return DB
}
