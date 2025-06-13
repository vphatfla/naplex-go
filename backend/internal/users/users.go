package users

import (
	"context"

	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

type UserController struct {
	queries *database.Queries
}

func NewUserController(queries *database.Queries) *UserController {
	return &UserController{
	 	queries: queries,
	}
}

func (c *UserController) GetUserInfo(ctx context.Context, uid int32) (*database.User, error) {
	u, err := c.queries.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &u, err
}
