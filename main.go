package main

import (
	"fmt"
	"godbapi/api"

	// "godbapi/database/migrations"
	"godbapi/model"
	"godbapi/utils"
)

const projectDirName = "godbapi/"

// init is invoked before main()
func init() {
	utils.LoadEnv(projectDirName)
}

func main() {
	// cwd, err := os.Getwd()
	// fmt.Println(err)
	// fmt.Println(cwd)

	// user := model.User{
	// 	id:   1,
	// 	name: "Test",
	// 	age:  25,
	// }

	var user model.User

	user.ID = 1
	user.Name = "shabbir"
	user.Age = 25

	// migrations.Create()

	obj := api.Sentiment()
	fmt.Println(obj)

	// fmt.Println(os.LookupEnv("db_host"))
}
