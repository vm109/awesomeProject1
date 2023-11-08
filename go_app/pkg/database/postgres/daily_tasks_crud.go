package postgres

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go_app/m/v2/pkg/domain/model"
	"time"
)

type dailyTask struct {
	Id             string     `db:"id"`
	TaskName       string     `db:"task_name"`
	CreatedAt      *time.Time `db:"created_at"`
	CompletedToday bool       `db:"completed_today"`
}

type runbookTask struct {
	Id          string     `db:"id"`
	TaskId      string     `db:"task_id"`
	CreatedAt   *time.Time `db:"created_at"`
	RunbookTask string     `db:"runbook_task"`
}

func newDailyTask(dt *model.DailyTask) *dailyTask {
	row := &dailyTask{
		Id:             dt.Id,
		TaskName:       dt.TaskName,
		CreatedAt:      dt.CreatedAt,
		CompletedToday: dt.CompletedToday,
	}
	return row
}
func (d *Dao) CreateDailyTask(dt *model.DailyTask) (*model.DailyTask, error) {
	row := newDailyTask(dt)
	err := d.db.QueryRowx(`INSERT INTO 
										dailytasks (task_name, created_at, completed_today)
								  VALUES
										($1, $2, $3)
								  RETURNING
										id, task_name, created_at, completed_today`,
		row.TaskName,
		row.CreatedAt,
		row.CompletedToday,
	).StructScan(row)

	e, ok := err.(*pq.Error)
	if ok && e.Code == "PS006" {
		return nil, errors.New("Duplicate Key - task already exists")
	} else if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dt.Id = row.Id
	return dt, nil
}
