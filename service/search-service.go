package service

import "qkeruen/repository"

type SearchService interface {
	Check(places string) (bool, error)
	Create(places string) error
	CheckGeo(places string) (bool, error)
	CreateGeo(places string) error
}

type searchService struct {
	db repository.SearchDB
}

func NewSearchService(ds repository.SearchDB) *searchService {
	return &searchService{db: ds}
}

func (s *searchService) Check(places string) (bool, error) {
	ok, err := s.db.Check(places)

	if err != nil {
		return false, nil
	}

	return ok, err
}

func (s *searchService) Create(places string) error {
	return s.db.Create(places)
}

func (s *searchService) CheckGeo(places string) (bool, error) {
	ok, err := s.db.CheckGeo(places)

	if err != nil {
		return false, nil
	}

	return ok, err
}

func (s *searchService) CreateGeo(places string) error {
	return s.db.CreateGeo(places)
}
