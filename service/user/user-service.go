package user

import "github.com/neostefan/diet-backend/models"

type UserService interface {
	GetUser(userId int) (*models.User, error)
}
