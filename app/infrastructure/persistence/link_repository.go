package persistence

import (
	"github.com/gregvroberts/hackernews/app/domain/repository/link"
	"github.com/jmoiron/sqlx"
)

type linkService struct {
	db *sqlx.DB
}

func NewLink(db *sqlx.DB) *userService {
	return &userService{
		db,
	}
}

var _ link.LinkService = &linkService{}
