package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neostefan/diet-backend/middlewares"
	"github.com/neostefan/diet-backend/models"
	"github.com/neostefan/diet-backend/service/planner"
	"github.com/neostefan/diet-backend/service/user"
	"github.com/neostefan/ga-diet/definitions"
)

type PlannerController struct {
	plannerService planner.PlannerService
	userService    user.UserService
}

func NewPlannerController(pS planner.PlannerService, uS user.UserService) PlannerController {
	return PlannerController{
		plannerService: pS,
		userService:    uS,
	}
}

func (pl *PlannerController) GetDashBoard(ctx *gin.Context) {
	var userId int
	var err error

	if userId, err = strconv.Atoi(ctx.Request.Header.Get("Userid")); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	u, err := pl.userService.GetUser(userId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	meals, err := pl.plannerService.GetDashBoard(userId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meals": meals, "user": u})

}

func (pl *PlannerController) CreateMealPlan(ctx *gin.Context) {
	var userConstraints *models.UserConstraints
	var err error
	var userId int
	var dietConditions []definitions.DietCondition

	if userId, err = strconv.Atoi(ctx.Request.Header.Get("Userid")); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	u, err := pl.userService.GetUser(userId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	switch u.Condition {
	case "diabetes":
		dietConditions = append(dietConditions, definitions.DIABETES)
	case "ulcer":
		dietConditions = append(dietConditions, definitions.ULCER)
	}

	if err := ctx.ShouldBindJSON(&userConstraints); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	meal, err := pl.plannerService.CreateMealPlan(userConstraints, dietConditions)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"meal": meal})
}

func (pl *PlannerController) RegisterPlannerRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/planner")

	r.GET("/d", middlewares.AuthMiddleware, pl.GetDashBoard)
	r.POST("/create", middlewares.AuthMiddleware, pl.CreateMealPlan)
}
