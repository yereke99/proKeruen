package dto

type OrderRequest struct {
	UserId        int64  `json:"userId"`
	LatitudeFrom  string `json:"latitudefrom"`
	LongitudeFrom string `json:"longitudeFrom"`
	LatitudeTo    string `json:"latitudeTo"`
	LongitudeTo   string `json:"longitudeTo"`
	Comments      string `json:"comments"`
	Price         int    `json:"price"`
	Type          string `json:"type"`
	OrderStatus   int    `json:"orderStatus"`
}

type OrderResponse struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`
	LatitudeFrom  string `json:"latitudefrom"`
	LongitudeFrom string `json:"longitudeFrom"`
	LatitudeTo    string `json:"latitudeTo"`
	LongitudeTo   string `json:"longitudeTo"`
	Comments      string `json:"comments"`
	Price         int    `json:"price"`
	Type          string `json:"type"`
	OrderStatus   int    `json:"orderStatus"`
}
