package model

import "time"

/*
Daily tasks are routines of the users.
A Collection of Daily Tasks would be a Routine
*/
type DailyTask struct {
	Id             string     `json:"id"`
	TaskName       string     `json:"task_name"`
	CreatedAt      *time.Time `json:"-"`
	CompletedToday bool       `json:"completed_today"`
}

type Routine struct {
	Id               string       `json:"id"`
	Tasks            []*DailyTask `json:"tasks"`
	CreatedAt        *time.Time   `json:"-"`
	CompletionStatus float32      `json:"completion_status"`
	RoutineName      string       `json:"routine_name"`
}
