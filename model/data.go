package model

import "time"

type Data struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Content   *string   `json:"content"`
	Key       *string   `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
