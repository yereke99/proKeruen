package service

import (
	"fmt"
	"log"
	"qkeruen/dto"
	"qkeruen/help"
	"qkeruen/repository"
)

type OrderService interface {
	CreateOrder(order dto.OrderRequest) error
	GetOrders(id int64, loc help.Location) ([]*dto.OrderResponse, error)
	GetMyOrders(id int64) ([]*dto.OrderResponse, error)
	DeleteOrder(orderId int64) error
}

type orderService struct {
	db repository.OrderDB
}

func NewOrderService(order repository.OrderDB) *orderService {
	return &orderService{db: order}
}

func (s *orderService) CreateOrder(order dto.OrderRequest) error {
	return s.db.CreateOrder(order)
}

func (s *orderService) GetOrders(id int64, loc help.Location) ([]*dto.OrderResponse, error) {
	var orders []*dto.OrderResponse

	res, err := s.db.GetOrders(id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for i, v := range res {
		log.Println(v.LatitudeFrom)
		log.Println(v.LongitudeFrom)
		if help.CheckDistance(loc, help.Convert(v.LatitudeFrom), help.Convert(v.LongitudeFrom)) {
			orders = append(orders, res[i])
		}
	}

	return orders, nil
}

func (s *orderService) GetMyOrders(id int64) ([]*dto.OrderResponse, error) {
	res, err := s.db.GetMyOrders(id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, err
}

func (s *orderService) DeleteOrder(orderId int64) error {
	return s.db.DeleteOrder(orderId)
}
