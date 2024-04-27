package model

import "time"

type Data struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Content    *string   `json:"content"`
	Key        *string   `json:"key"`
	PublicKey  *string   `json:"public_key"`
	PrivateKey *string   `json:"private_key"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}
