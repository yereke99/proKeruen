package models

type History struct {
	Id           int64 `json:"id"`
	OrderId      int64 `json:"orderId"`
	DriverId     int64 `json:"driverId"`
	UserId       int64 `json:"userId"`
	StartDate    string   `json:"startDate"`
	FinishedDate string   `json:"finishedDate"`
}
