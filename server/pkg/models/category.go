package models

type Category struct {
	Name string `json:"name" gorm:"unique; not null;"`
	BaseModel
}
