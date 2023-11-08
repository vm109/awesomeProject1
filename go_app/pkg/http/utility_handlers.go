package http

import (
	"net/http"
)

func (a *Api) statushandler(w http.ResponseWriter, r *http.Request) {
	type statusMessage struct {
		AppName string
		Status  string
	}
	body := statusMessage{AppName: "TaskManager", Status: "Live"}

	a.WriteJsonBodyResponse(r, w, body, http.StatusOK)
}
