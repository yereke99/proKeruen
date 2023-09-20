package service

import (
	"qkeruen/models"
	"qkeruen/repository"
)

type HistoryService interface {
	GetUserHistory(id int64) ([]*models.History, error)
	GetDriverHistory(id int64) ([]*models.History, error)
}

type historyService struct {
	db repository.HistoryDB
}

func NewHistoryService(ds repository.HistoryDB) *historyService {
	return &historyService{
		db: ds,
	}
}

func (s *historyService) GetUserHistory(id int64) ([]*models.History, error) {
	res, err := s.db.GetUserHistory(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *historyService) GetDriverHistory(id int64) ([]*models.History, error) {
	res, err := s.db.GetDriverHistory(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
