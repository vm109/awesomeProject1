package http

import (
	"encoding/json"
	"fmt"
	"go_app/m/v2/pkg/database/postgres"
	"go_app/m/v2/pkg/domain/model"
	"io"
	"net/http"
	"time"
)

func supaClient() *postgres.Dao {
	return postgres.CreateSupaClient()
}

func (a *Api) createASingleDailyTask(w http.ResponseWriter, r *http.Request) {
	var requestBody []byte
	var err error
	var dailyTask model.DailyTask
	fmt.Println(r.URL)
	if requestBody, err = io.ReadAll(r.Body); err != nil {
		a.WriteErrorResponse(r, w, err)
	}

	if requestBody != nil && len(requestBody) > 0 {
		fmt.Println(requestBody)
		if err = json.Unmarshal(requestBody, &dailyTask); err != nil {
			fmt.Println(err)
			a.WriteErrorResponse(r, w, err)
		} else {
			if dailyTask.CreatedAt == nil {
				now := time.Now().UTC()
				dailyTask.CreatedAt = &now
			}
			if dt, err := supaClient().CreateDailyTask(&dailyTask); err != nil {
				a.WriteErrorResponse(r, w, err)
			} else {
				a.WriteJsonBodyResponse(r, w, dt, http.StatusOK)
			}
		}
	} else {
		err = model.InvalidData()
		a.WriteErrorResponse(r, w, err)
	}
}
