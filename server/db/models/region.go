package models

type Region struct {
	Name string `json:"name" gorm:"unique; not null;"`
	BaseModel
}
