package models

type OfferUserModel struct {
	Id        int64  `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Price     int    `json:"price"`
	Comment   string `json:"comment"`
	Type      string `json:"type"`
	Driver    int64  `json:"driver"`
	DriverAVA string `json:"ava"`
}
