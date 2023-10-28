package handlers

import (
	"encoding/json"
	"go_app/m/v2/pkg/http"
	"net/http"
)

func (a *Api) statushandler(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal("hello")

	a.WriteJsonBodyResponse(r, w, body, http.StatusOK)
}
