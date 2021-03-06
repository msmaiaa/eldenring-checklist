package models

type Check struct {
	UserId   string `json:"userId" gorm:"not null"`
	User     User   `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	EntityId uint   `json:"entityId" gorm:"not null"`
	Entity   Entity `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}
