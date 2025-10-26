package service

import "crud-backend/usuarios/entity"

type UsuarioService interface {
	FindAll() ([]entity.Usuario, error)
	FindByID(id uint) (*entity.Usuario, error)
	Save(usuario *entity.Usuario) (*entity.Usuario, error)
	Update(usuario *entity.Usuario) (*entity.Usuario, error)
	Delete(id uint) error
}
