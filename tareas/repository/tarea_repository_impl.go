package repository

import (
	"crud-backend/database"
	"crud-backend/tareas/entity"
)

type tareaRepositoryImpl struct{}

func NewTareaRepository() TareaRepository {
	return &tareaRepositoryImpl{}
}

func (r *tareaRepositoryImpl) FindAll() ([]entity.Tarea, error) {
	var tareas []entity.Tarea
	result := database.DB.Find(&tareas)
	return tareas, result.Error
}

func (r *tareaRepositoryImpl) FindByID(id uint) (*entity.Tarea, error) {
	var tarea entity.Tarea
	result := database.DB.First(&tarea, id)
	return &tarea, result.Error
}

func (r *tareaRepositoryImpl) Save(t *entity.Tarea) (*entity.Tarea, error) {
	result := database.DB.Create(t)
	return t, result.Error
}

func (r *tareaRepositoryImpl) Update(t *entity.Tarea) (*entity.Tarea, error) {
	result := database.DB.Save(t)
	return t, result.Error
}

func (r *tareaRepositoryImpl) Delete(id uint) error {
	result := database.DB.Delete(&entity.Tarea{}, id)
	return result.Error
}
