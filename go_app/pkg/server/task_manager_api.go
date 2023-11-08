package server

import (
	"go_app/m/v2/pkg/http"
)

type TaskManagerServer struct {
}

func (taskManager *TaskManagerServer) Boot() error {
	api := http.NewTaskManagerRestApi(8080)
	api.Start()
	return nil
}
