package entity

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string `gorm:"type:varchar(100);not null" json:"nombre"`
	Correo   string `gorm:"type:varchar(100);unique;not null" json:"correo"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}
