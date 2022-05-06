package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Steamid64 string `json:"steamid64"`
}

type Region struct {
	gorm.Model
	Name string `json:"name"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Entity struct {
	gorm.Model
	Id int8 `json:"id"`
}

type Coordinate struct {
	gorm.Model
	X int16 `json:"x"`
	Y int16 `json:"y"`
}