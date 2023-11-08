package postgres

import (
	"encoding/json"
	"fmt"
	"go_app/m/v2/pkg/domain/model"
	"time"
)

type plan struct {
	Id              string          `db:"id"`
	PlanDescription json.RawMessage `db:"plan_description"`
	CreatedAt       *time.Time      `db:"created_at"`
	UpdatedAt       *time.Time      `db:"updated_at"`
	ProgressPhase   *string         `db:"progress_phase"`
}

func newPlan() {

}

func (p *plan) Model() (planModel *model.Plan) {

	planModel = &model.Plan{
		Id:              p.Id,
		CreatedDate:     p.CreatedAt,
		UpdatedDate:     p.UpdatedAt,
		PlanDescription: p.PlanDescription,
	}
	if p.ProgressPhase != nil {
		planModel.ProgressPhase = *p.ProgressPhase
	}
	return
}

func (d *Dao) GetPlanById(plan_id string) (planModel *model.Plan, err error) {
	var dbPlan plan
	query := fmt.Sprintf(
		`SELECT id, plan_description, created_at, updated_at, progress_phase FROM plans where id='%s'`,
		plan_id,
	)
	err = d.Get(&dbPlan, query)
	if err != nil {
		return nil, err
	} else {
		planModel = dbPlan.Model()
		return
	}
}
