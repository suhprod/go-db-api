package database

import (
	"godbapi/model"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//WriteUser add user to db
func WriteUser(user model.User) {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

}
