package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (a *Api) getPlanById(w http.ResponseWriter, r *http.Request) {
	planId := httprouter.ParamsFromContext(r.Context()).ByName("id")
	plan, err := supaClient().GetPlanById(planId)
	if err != nil {
		fmt.Println(err)
		a.WriteErrorResponse(r, w, err)
	} else {
		a.WriteJsonBodyResponse(r, w, plan, http.StatusOK)
	}
}
