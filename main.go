package main

import (
	"fmt"
	"todo_app_heroku/app/controllers"
	"todo_app_heroku/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
