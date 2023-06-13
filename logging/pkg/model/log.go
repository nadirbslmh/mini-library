package model

import "time"

type Log struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	UserID    int        `json:"user_id"`
	BookID    int        `json:"book_id"`
	BookTitle string     `json:"book_title"`
	CreatedAt time.Time  `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}
