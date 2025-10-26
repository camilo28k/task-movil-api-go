package repository

import (
	"crud-backend/database"
	"crud-backend/usuarios/entity"
)

type usuarioRepositoryImpl struct{}

func NewUsuarioRepository() UsuarioRepository {
	return &usuarioRepositoryImpl{}
}

func (r *usuarioRepositoryImpl) FindAll() ([]entity.Usuario, error) {
	var usuarios []entity.Usuario
	result := database.DB.Find(&usuarios)
	return usuarios, result.Error
}

func (r *usuarioRepositoryImpl) FindByID(id uint) (*entity.Usuario, error) {
	var usuario entity.Usuario
	result := database.DB.First(&usuario, id)
	return &usuario, result.Error
}

func (r *usuarioRepositoryImpl) Save(usuario *entity.Usuario) (*entity.Usuario, error) {
	result := database.DB.Create(usuario)
	return usuario, result.Error
}

func (r *usuarioRepositoryImpl) Update(usuario *entity.Usuario) (*entity.Usuario, error) {
	result := database.DB.Save(usuario)
	return usuario, result.Error
}

func (r *usuarioRepositoryImpl) Delete(id uint) error {
	result := database.DB.Delete(&entity.Usuario{}, id)
	return result.Error
}
