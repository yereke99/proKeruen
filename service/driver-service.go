package service

import (
	"fmt"
	"qkeruen/models"
	"qkeruen/repository"
)

type DriverService interface {
	CreateDriver(data models.DriverRegister) (*models.DriverModel, error)
	GetProfile(token string) (*models.DriverModel, error)
	UpdateService(update models.DriverModel) (*models.DriverModel, error)
	Delete(id int64) error
}

type driverService struct {
	db repository.DriverDB
}

func NewDriverService(ds repository.DriverDB) *driverService {
	return &driverService{
		db: ds,
	}
}

func (s *driverService) CreateDriver(data models.DriverRegister) (*models.DriverModel, error) {
	driver, err := s.db.InsertDriverData(data)

	if err != nil {
		return driver, err
	}

	return driver, nil
}

func (s *driverService) GetProfile(token string) (*models.DriverModel, error) {
	data, err := s.db.GetDriverProfile(token)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *driverService) UpdateService(update models.DriverModel) (*models.DriverModel, error) {
	res, err := s.db.UpdateDriver(update)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *driverService) Delete(id int64) error {
	return s.db.Delete(id)
}
