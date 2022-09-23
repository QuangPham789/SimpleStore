package entity

type NewOrder struct {
	Accountid int     `json:"accountid"`
	Items     []*Item `json:"items"`
	Total     int     `json:"total"`
}

type Order struct {
	ID        string  `json:"id"`
	Accountid int     `json:"accountid"`
	Items     []*Item `json:"items"`
	Itemid    int     `json:"itemid"`
	Total     int     `json:"total"`
}

type UpdateOrderModel struct {
	ID    int     `json:"id"`
	Items []*Item `json:"items"`
	Total int     `json:"total"`
}
