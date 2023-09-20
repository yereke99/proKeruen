package models

type UserModel struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Token     string `josn:"token"`
}

type UserModelForDriver struct {
	Id        int64  `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
}
