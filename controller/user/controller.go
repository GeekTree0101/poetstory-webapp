package user

import (
	context "context"

	"github.com/GeekTree0101/poetstory-webapp/domain/entity"
	"github.com/GeekTree0101/poetstory-webapp/domain/usecase"
)

func New(fetchUserUsecase *usecase.FetchUsersUsecase) *Controller {
	return &Controller{
		fetchUserUsecase: fetchUserUsecase,
	}
}

// Controller for users
type Controller struct {
	fetchUserUsecase *usecase.FetchUsersUsecase
}

// Index of users
// GET /user
func (c *Controller) Index(ctx context.Context) (users []*entity.User, err error) {
	users, err = c.fetchUserUsecase.Users()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Show user
// GET /user/:id
func (c *Controller) Show(ctx context.Context, id int) (user *entity.User, err error) {
	user, err = c.fetchUserUsecase.UserByID(uint(id))

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete user
// DELETE /user/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
