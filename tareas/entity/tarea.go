package entity

import "gorm.io/gorm"

type Tarea struct {
	gorm.Model
	Titulo      string `gorm:"type:varchar(100);not null" json:"titulo"`
	Descripcion string `gorm:"type:text" json:"descripcion"`
	Completada  bool   `gorm:"default:false" json:"completada"`
	UsuarioID   uint   `json:"usuario_id"` // relación con Usuario
}
