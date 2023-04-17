package planner

import (
	"github.com/neostefan/diet-backend/models"
	"github.com/neostefan/ga-diet/definitions"
)

type PlannerService interface {
	CreateMealPlan(uc *models.UserConstraints, conditions []definitions.DietCondition) (*models.Meal, error)
}
