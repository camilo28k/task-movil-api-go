package dto

type TareaRequest struct {
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Completada  bool   `json:"completada"`
	UsuarioID   uint   `json:"usuario_id"`
}
