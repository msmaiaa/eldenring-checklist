package models

import "time"

type User struct {
	Steamid64 string    `json:"steamid64" gorm:"primaryKey;unique;not null;"`
	Role      string    `json:"role" gorm:"default:user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}