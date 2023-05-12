package main

import (
	"rohan.com/task-manager/cmd/task-manager/app"
)

func main() {
	a := app.New()
	err := a.Start()
	if err != nil {
		print("could not start application. Error:: ", err.Error())
	}

	defer func() {
		if err := a.Shutdown(); err != nil {
			panic(err)
		}
	}()
}
