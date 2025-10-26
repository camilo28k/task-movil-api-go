package service

import (
	"crud-backend/tareas/entity"
	"crud-backend/tareas/repository"
)

type tareaServiceImpl struct {
	repo repository.TareaRepository
}

func NewTareaService(r repository.TareaRepository) TareaService {
	return &tareaServiceImpl{repo: r}
}

func (s *tareaServiceImpl) FindAll() ([]entity.Tarea, error) {
	return s.repo.FindAll()
}

func (s *tareaServiceImpl) FindByID(id uint) (*entity.Tarea, error) {
	return s.repo.FindByID(id)
}

func (s *tareaServiceImpl) Save(t *entity.Tarea) (*entity.Tarea, error) {
	return s.repo.Save(t)
}

func (s *tareaServiceImpl) Update(t *entity.Tarea) (*entity.Tarea, error) {
	return s.repo.Update(t)
}

func (s *tareaServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
