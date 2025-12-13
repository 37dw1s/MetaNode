package main

import (
	"GoLearning/task4/internal/database"
	"GoLearning/task4/internal/router"
)

func main() {
	err := database.Init()
	if err != nil {
		return
	}

	r := router.Setup()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
