package models

type ProcessModel struct {
	Id        int64 `json:"id"`
	OrderId   int64 `json:"orderId"`
	DriverId  int64 `json:"driverId"`
	UserId    int64 `json:"userId"`
	StartDate int64 `json:"startDate"`
}
