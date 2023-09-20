package models

import "database/sql"

type Security struct {
	Id         int64          `json:"id"`
	UserId     int64          `json:"userId"`
	FirsrtName string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	From       string         `json:"from"`
	To         string         `json:"to"`
	FioD       string         `json:"fioD"`
	Phone      string         `json:"phone"`
	CarNumber  string         `json:"carNumber"`
	TimeStart  string         `json:"timeStart"`
	TimeFinish sql.NullString `json:"timeFinish"`
	Check      string         `json:"check"`
}
