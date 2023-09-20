package models

type OfferDriverModel struct {
	Id      int64  `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Price   int    `json:"price"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
	User    int64  `json:"user"`
	UserAVA string `json:"ava"`
}
