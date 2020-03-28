package migrations

import "github.com/jinzhu/gorm"

//users struct
type users struct {
	ID   int64
	Name string `sql:"size:255"`
	Age  int64
}

var db gorm.DB

//Create function
func Create() {
	db.CreateTable(&users{})
}

//Remove function
func Remove() {
	db.DropTableIfExists(&users{}, "users")
}
