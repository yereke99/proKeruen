package service

import (
	"qkeruen/models"
	"qkeruen/repository"
)

type UserService interface {
	Create(data models.UserRegister) (*models.UserModel, error)
	CheckTokenUser(token string) (*models.UserModel, error)
	Update(user models.UserModel) (*models.UserModel, error)
	Delete(id int) error
}

type userService struct {
	db repository.UserDB
}

func NewUserService(db repository.UserDB) *userService {
	return &userService{db: db}
}

func (s *userService) Create(data models.UserRegister) (*models.UserModel, error) {
	user, err := s.db.Insert(data)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CheckTokenUser(token string) (*models.UserModel, error) {
	res, err := s.db.CheckTokenUser(token)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) Update(user models.UserModel) (*models.UserModel, error) {
	res, err := s.db.Update(user)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *userService) Delete(id int) error {
	return s.db.Delete(id)
}
