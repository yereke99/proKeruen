package dto

type OfferRequest struct {
	//Id      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Price   int    `json:"price"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
}
