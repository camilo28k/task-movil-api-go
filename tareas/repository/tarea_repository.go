package repository

import "crud-backend/tareas/entity"

type TareaRepository interface {
	FindAll() ([]entity.Tarea, error)
	FindByID(id uint) (*entity.Tarea, error)
	Save(tarea *entity.Tarea) (*entity.Tarea, error)
	Update(tarea *entity.Tarea) (*entity.Tarea, error)
	Delete(id uint) error
}
