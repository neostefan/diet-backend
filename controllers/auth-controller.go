package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neostefan/diet-backend/models"
	"github.com/neostefan/diet-backend/service/auth"
	"github.com/neostefan/diet-backend/util"
)

type AuthController struct {
	authService auth.AuthenticationService
}

func NewAuthController(authS auth.AuthenticationService) AuthController {
	return AuthController{
		authService: authS,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ac.authService.Register(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": true})
}

func (ac *AuthController) SignIn(ctx *gin.Context) {
	var credentials models.SignInCredentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := ac.authService.SignIn(credentials.FirstName, credentials.LastName, credentials.Password)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	tokenString, err := util.CreateJWT(int(u.Id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": true, "token": tokenString})
}

func (a *AuthController) RegisterAuthenticationRoutes(rg *gin.RouterGroup) {
	rg = rg.Group("/auth")
	rg.POST("/register", a.Register)
	rg.POST("/sign-in", a.SignIn)
}
