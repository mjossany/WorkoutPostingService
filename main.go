package main

import (
	"github.com/mjossany/workoutPostingService/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("we are running our app")
}
