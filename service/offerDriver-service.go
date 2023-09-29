package service

import (
	"qkeruen/dto"
	"qkeruen/models"
	"qkeruen/repository"
)

type OfferDriverService interface {
	GetByID(id int64) (*models.UserModelForDriver, error)
	CreateOffer(id int, offer dto.OfferRequest) error
	MyOffer(id int64) ([]*dto.OfferResponseDriver, error)
	FindAllOffers() ([]*models.OfferDriverModel, error)
	SearchOffers(to, from, type_ string) ([]*models.OfferDriverModel, error)
	Delete(offerId int64) error
}

type offerDriverService struct {
	db repository.OfferDriverDB
}

func NewOfferDriverService(ds repository.OfferDriverDB) *offerDriverService {
	return &offerDriverService{
		db: ds,
	}
}

func (s *offerDriverService) CreateOffer(id int, offer dto.OfferRequest) error {
	return s.db.Create(id, offer)
}

func (s *offerDriverService) GetByID(id int64) (*models.UserModelForDriver, error) {
	res, err := s.db.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *offerDriverService) MyOffer(id int64) ([]*dto.OfferResponseDriver, error) {
	res, err := s.db.MyOffer(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *offerDriverService) FindAllOffers() ([]*models.OfferDriverModel, error) {
	res, err := s.db.FindAllOffers()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *offerDriverService) SearchOffers(to, from, type_ string) ([]*models.OfferDriverModel, error) {
	if to == "" {
		res, err := s.db.SearchOneSide(from, type_)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
	res, err := s.db.Search(to, from, type_)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *offerDriverService) Delete(offerId int64) error {
	return s.db.Delete(offerId)
}
