package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID           int `gorm:"primaryKey"`
	Title        string
	ReleasedYear string `json:"releasedyear"`
	Raiting      string `json:"raiting"`
	IDMovie      string `json:"idmovie"`
	Genre        string `json:"genre"`
}
