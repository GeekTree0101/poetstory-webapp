package repository

import "github.com/GeekTree0101/poetstory-webapp/domain/entity"

type UserRepository struct {
}

func DefaultUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetUsers() ([]*entity.User, error) {
	// dummy
	return []*entity.User{
		{
			Nickname: "David",
			Bio:      "Hello~ i'm a sofeware engineer",
		},
		{
			Nickname: "Claire",
			Bio:      "Niha~ I'm a data analyst",
		},
	}, nil
}
