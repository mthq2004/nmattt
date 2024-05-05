package model

import "time"

type Data struct {
	ID                string    `json:"id" gorm:"primaryKey"`
	Type              string    `json:"type"`
	Encrypted_content string    `json:"encrypted_content"`
	Content           *string   `json:"content"`
	Key               *string   `json:"key"`
	PublicKey         *string   `json:"public_key"`
	PrivateKey        *string   `json:"private_key"`
	CreatedAt         time.Time `json:"created_at"`
}
