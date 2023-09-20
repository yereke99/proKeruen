package service

import (
	"qkeruen/models"
	"qkeruen/repository"
	"time"
)

type SecurityService interface {
	Create(data models.Security) error
	GetMyHistory(id int64) ([]*models.Security, error)
	Finish(id int64) (*models.Security, error)
}

type securityService struct {
	db repository.SecurityDB
}

func NewSecurityService(db repository.SecurityDB) *securityService {
	service := &securityService{
		db: db,
	}

	return service
}

func (s *securityService) Create(data models.Security) error {
	currentTime := time.Now()
	timeNow := currentTime.Format("2006-01-02 15:04:05")
	data.TimeStart = timeNow
	data.Check = "false"

	if err := s.db.Insert(data); err != nil {
		return err
	}

	return nil
}

func (s *securityService) GetMyHistory(id int64) ([]*models.Security, error) {
	res, err := s.db.GetMyHistory(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *securityService) Finish(id int64) (*models.Security, error) {
	timeNow := time.Now()
	timeNowStr := timeNow.Format(time.RFC3339)

	res, err := s.db.Finish(id, timeNowStr)
	if err != nil {
		return nil, err
	}

	return res, nil
}
