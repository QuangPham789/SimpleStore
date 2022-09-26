package entity

type NewOrder struct {
	Accountid int     `json:"accountid"`
	Items     []*Item `json:"items"`
	Total     int     `json:"total"`
}

type Order struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Accountid  int    `json:"accountid"`
	TotalPrice int    `json:"totalprice"`
}

type UpdateOrderModel struct {
	ID    int     `json:"id"`
	Items []*Item `json:"items"`
	Total int     `json:"total"`
}
