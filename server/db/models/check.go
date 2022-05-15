package models

type Check struct {
	UserId uint `json:"userId" gorm:"not null"`
	EntityId uint `json:"entityId" gorm:"not null"`
}