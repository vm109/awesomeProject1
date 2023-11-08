package main

import (
	"fmt"
	"go_app/m/v2/pkg/server"
)

func main() {
	fmt.Print("Starting Personal Task Manager")
	taskManager := &server.TaskManagerServer{}
	if err := taskManager.Boot(); err != nil {
		panic(err)
	}
}
