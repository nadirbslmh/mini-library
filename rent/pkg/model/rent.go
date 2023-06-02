package model

type Rent struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
