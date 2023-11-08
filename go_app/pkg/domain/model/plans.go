package model

import (
	"encoding/json"
	"time"
)

type Plan struct {
	Id              string          `json:"id"`
	PlanDescription json.RawMessage `json:"plan_description"`
	CreatedDate     *time.Time      `json:"created_date"`
	UpdatedDate     *time.Time      `json:"updated_date"`
	ProgressPhase   string          `json:"progress_phase"`
}
