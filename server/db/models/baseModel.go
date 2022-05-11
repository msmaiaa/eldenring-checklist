package models

import (
	"time"
)

type BaseModel struct {
	Id uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}