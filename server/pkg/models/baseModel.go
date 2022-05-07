package models

import (
	"time"
)

type BaseModel struct {
	Id uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time 
  UpdatedAt time.Time 
}