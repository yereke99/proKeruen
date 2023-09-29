package models

type DriverModel struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Inn       string `json:"inn"`
	Avatar    string `json:"avatar"`
	CarNumber string `json:"carNumber"`
	CarColor  string `json:"carColor"`
	CarModel  string `json:"carModel"`
	DocsFront string `json:"docsfront"`
	DocsBacks string `json:"docsback"`
	CarType   string `json:"carType"`
	Token     string `json:"token"`
}

type DriverModelForUser struct {
	Id        int64  `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	//Inn       string `json:"inn"`
	Avatar    string `json:"avatar"`
	CarNumber string `json:"carNumber"`
	CarColor  string `json:"carColor"`
	CarModel  string `json:"carModel"`
	//DocsFront string `json:"docsfront"`
	//DocsBacks string `json:"docsback"`
}
