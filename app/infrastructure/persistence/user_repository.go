package persistence

import (
	"github.com/gregvroberts/hackernews/app/domain/repository/user"
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *userService {
	return &userService{
		db,
	}
}

var _ user.UserService = &userService{}
