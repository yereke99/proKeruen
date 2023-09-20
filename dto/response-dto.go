package dto

type ResponseDTO struct {
	Token        string `json:"token"`
	IsAuthorized bool   `json:"isAuthorized"`
}
