package auth

import (
	"github.com/neostefan/diet-backend/models"
)

type AuthenticationService interface {
	Register(u *models.User) error
	SignIn(firstname, lastname, password string) (models.User, error)
}
