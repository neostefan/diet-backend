package planner

import (
	"context"
	"database/sql"

	"github.com/neostefan/diet-backend/db"
	"github.com/neostefan/diet-backend/models"
	nsga "github.com/neostefan/ga-diet"
	"github.com/neostefan/ga-diet/definitions"
)

type PlannerServiceImpl struct {
	database *sql.DB
	ctx      context.Context
}

func New(db *sql.DB, ctx context.Context) PlannerServiceImpl {
	return PlannerServiceImpl{
		database: db,
		ctx:      ctx,
	}
}

func (pS PlannerServiceImpl) CreateMealPlan(uc *models.UserConstraints, conditions []definitions.DietCondition) (*models.Meal, error) {
	var meal models.Meal
	totalCalories := 0.0
	totalCost := 0.0
	totalProtein := 0.0

	ings := nsga.Nsga(uc.Max, uc.Min, conditions)

	for _, ing := range ings {
		if ing.Type == "carbs" {
			meal.Carbs = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}

		if ing.Type == "protein" {
			meal.Proteins = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}

		if ing.Type == "fruits" {
			meal.Fruits = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}

		if ing.Type == "oils" {
			meal.Oils = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}

		if ing.Type == "beverages" {
			meal.Beverages = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}

		if ing.Type == "vegetables" {
			meal.Vegetables = ing.Name
			totalCalories += ing.Calories
			totalCost += ing.Cost
			totalProtein += ing.Protein
		}
	}

	meal.Calories = totalCalories
	meal.Cost = totalCost
	meal.ProteinValue = totalProtein

	//Add meal plan to the recommendation table
	err := db.AddRecommendation(pS.database, &meal)

	if err != nil {
		return nil, err
	}

	return &meal, nil
}
