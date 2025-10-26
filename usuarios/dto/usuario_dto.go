package dto

type UsuarioRequest struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type UsuarioResponse struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}
