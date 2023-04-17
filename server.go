package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/neostefan/diet-backend/controllers"
	"github.com/neostefan/diet-backend/db"
	"github.com/neostefan/diet-backend/middlewares"
	"github.com/neostefan/diet-backend/service/auth"
	"github.com/neostefan/diet-backend/service/planner"
	"github.com/neostefan/diet-backend/service/user"
)

var (
	authService       auth.AuthenticationService
	plannerService    planner.PlannerService
	userService       user.UserService
	authController    controllers.AuthController
	plannerController controllers.PlannerController
	backend_Db        *sql.DB
	err               error
)

func init() {
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)
	ctx := context.TODO()

	backend_Db, err = sql.Open("sqlite3", "./db/backend.db")

	if err != nil {
		log.Fatalf("\n Unable to Open database: %s \n", err.Error())
	}

	db.CreateUsersTable(backend_Db)
	db.CreateRecommendationsTable(backend_Db)

	authService = auth.New(backend_Db, ctx)
	plannerService = planner.New(backend_Db, ctx)
	userService = user.New(backend_Db, ctx)
	authController = controllers.NewAuthController(authService)
	plannerController = controllers.NewPlannerController(plannerService, userService)
}

func main() {
	s := gin.New()

	s.Use(gin.LoggerWithFormatter(middlewares.Logger), gin.Recovery())

	basepath := s.Group("/api")

	authController.RegisterAuthenticationRoutes(basepath)
	plannerController.RegisterPlannerRoutes(basepath)

	s.Run(":5000")

	defer backend_Db.Close()
}
