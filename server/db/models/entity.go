package models

type Entity struct {
	Name        string   `json:"name" gorm:"unique; not null;"`
	X           int16    `json:"x" gorm:"not null"`
	Y           int16    `json:"y" gorm:"not null"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"-"`
	Category    Category `json:"category"`
	RegionID    uint     `json:"-"`
	Region      Region   `json:"region"`
	BaseModel
}