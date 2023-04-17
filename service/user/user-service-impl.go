package user

import (
	"context"
	"database/sql"

	"github.com/neostefan/diet-backend/db"
	"github.com/neostefan/diet-backend/models"
)

type UserServiceImpl struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) UserServiceImpl {
	return UserServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (us UserServiceImpl) GetUser(userId int) (*models.User, error) {
	user, err := db.GetUserById(us.db, userId)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
