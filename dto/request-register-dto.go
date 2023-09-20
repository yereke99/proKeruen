package dto

type RequestRegisterDTO struct {
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}

type CheckCodeRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Code        string `json:"code"`
	Role        string `json:"role"`
}
