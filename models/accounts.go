package models

type Accounts struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  int    `json:"password"`
	CreatedAt string `json:"created_at"`
}
