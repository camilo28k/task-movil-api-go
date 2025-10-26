package repository

import "crud-backend/usuarios/entity"

type UsuarioRepository interface {
	FindAll() ([]entity.Usuario, error)
	FindByID(id uint) (*entity.Usuario, error)
	Save(usuario *entity.Usuario) (*entity.Usuario, error)
	Update(usuario *entity.Usuario) (*entity.Usuario, error)
	Delete(id uint) error
}
