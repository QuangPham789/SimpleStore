package entity

type Item struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Productid int    `json:"productid"`
	Unit      int    `json:"unit"`
	Price     string `json:"price"`
	Total     string `json:"total"`
}
