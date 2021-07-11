package interfaces

import (
	"github.com/gregvroberts/hackernews/app/domain/repository/link"
	"github.com/gregvroberts/hackernews/app/domain/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	LinkService link.LinkService
	UserService user.UserService
}
