package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/neostefan/diet-backend/db"
	"github.com/neostefan/diet-backend/models"
)

type AuthServiceImpl struct {
	database *sql.DB
	ctx      context.Context //this is here for creating things like request limits
}

func New(database *sql.DB, ctx context.Context) AuthServiceImpl {
	return AuthServiceImpl{
		database: database,
		ctx:      ctx,
	}
}

func (au AuthServiceImpl) Register(u *models.User) error {
	var existingUser models.User
	var err error

	existingUser, err = db.GetUser(au.database, u.FirstName, u.LastName)

	if err != nil {
		return err
	}

	if (existingUser != models.User{}) {
		userErr := errors.New("user with that firstname and lastname already exists")
		return userErr
	}

	err = db.AddUser(au.database, u)

	if err != nil {
		return err
	}

	return nil
}

func (au AuthServiceImpl) SignIn(firstname string, lastname string, password string) (models.User, error) {
	var existingUser models.User
	var err error

	existingUser, err = db.GetUser(au.database, firstname, lastname)

	if err != nil {
		return existingUser, err
	}

	if (existingUser == models.User{}) {
		userErr := errors.New("no user with that firstname and lastname exists")
		return existingUser, userErr
	}

	if existingUser.Password != password {
		userErr := errors.New("invalid password")
		return existingUser, userErr
	}

	return existingUser, nil
}
