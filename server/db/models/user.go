package models

type User struct {
	Steamid64 string `json:"steamid64" gorm:"primaryKey; unique; not null;"`
	BaseModel
}