package service

import (
	"errors"

	"io/ioutil"
	"log"
	"net/http"
	"qkeruen/config"
	"qkeruen/models"
	"qkeruen/repository"
)

type AuthService interface {
	Check(contact, role string) (bool, error)
	GiveTokenService(contact, role string) (string, error)
	Create(contact string, code int) error
	ValidateSMS(contact string, code_validate int) (bool, error)
	CheckTokenDriver(tokens string) (*models.DriverModel, error)
	CheckTokenUser(token string) (*models.UserModel, error)
}

type authService struct {
	db repository.Database
}

func NewAuthService(ds repository.Database) *authService {
	return &authService{
		db: ds,
	}
}

func (s *authService) Check(contact, role string) (bool, error) {
	res, err := s.db.CheckFromRepo(contact, role)

	if err != nil {
		return false, err
	}

	if !res {
		return false, nil
	}

	return true, nil
}

func (s *authService) GiveTokenService(contact, role string) (string, error) {
	token, err := s.db.GiveToken(contact, role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Create(contact string, code int) error {
	e := s.db.CreateCode(contact, code)
	if e != nil {
		return e
	}

	resp, err := http.Get(config.ConfigSMS(contact, code))

	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

	return nil
}

func (s *authService) ValidateSMS(contact string, code_validate int) (bool, error) {
	s.db.All()

	code, err := s.db.ValidateSMS(contact)

	if err != nil {
		return false, err
	}

	if code == 0 && err != nil {
		return false, errors.New("The code id 0.")
	}

	if code_validate != code {
		return false, nil
	}

	return true, nil
}

// tommorow you must to change
func (s *authService) CheckTokenDriver(token string) (*models.DriverModel, error) {

	data, err := s.db.CheckTokenDriver(token)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *authService) CheckTokenUser(token string) (*models.UserModel, error) {

	data, err := s.db.CheckTokenUser(token)

	if err != nil {
		return nil, err
	}

	return data, nil
}
