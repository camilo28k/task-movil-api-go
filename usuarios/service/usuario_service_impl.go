package service

import (
	"crud-backend/usuarios/entity"
	"crud-backend/usuarios/repository"
)

type usuarioServiceImpl struct {
	repo repository.UsuarioRepository
}

func NewUsuarioService(r repository.UsuarioRepository) UsuarioService {
	return &usuarioServiceImpl{repo: r}
}

func (s *usuarioServiceImpl) FindAll() ([]entity.Usuario, error) {
	return s.repo.FindAll()
}

func (s *usuarioServiceImpl) FindByID(id uint) (*entity.Usuario, error) {
	return s.repo.FindByID(id)
}

func (s *usuarioServiceImpl) Save(u *entity.Usuario) (*entity.Usuario, error) {
	return s.repo.Save(u)
}

func (s *usuarioServiceImpl) Update(u *entity.Usuario) (*entity.Usuario, error) {
	return s.repo.Update(u)
}

func (s *usuarioServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
