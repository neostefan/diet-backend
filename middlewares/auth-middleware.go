package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	var userId string

	tokenString := strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1]
	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unidentified signing method: %v", t.Header["alg"])
		}

		return []byte("The cow says moo!!"), nil
	})

	if err != nil {
		fmt.Println("I am here o")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//! When exp time is given check if the current time has exceeded the expiration time
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//first convert the int to float64 since the type isn't a string initially
		uid := claims["id"].(float64)

		//convert the float64 to string
		userId = fmt.Sprintf("%.f", uid)

		ctx.Request.Header.Set("Userid", userId)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.Next()
}
