package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	err := controllers.StartMainServer()
	if err != nil {
		fmt.Println(err)
	}
}
