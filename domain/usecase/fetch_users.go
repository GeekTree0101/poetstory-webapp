package usecase

import (
	"errors"

	"github.com/GeekTree0101/poetstory-webapp/data/repository"
	"github.com/GeekTree0101/poetstory-webapp/domain/entity"
)

type FetchUsersUsecase struct {
	userRepository *repository.UserRepository
}

func DefaultFetchUsersUseCase() *FetchUsersUsecase {
	return &FetchUsersUsecase{
		userRepository: repository.DefaultUserRepository(),
	}
}

func (uc *FetchUsersUsecase) Users() ([]*entity.User, error) {
	users, err := uc.userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *FetchUsersUsecase) UserByID(id uint) (*entity.User, error) {
	users, err := uc.userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, errors.New("not found")
}
