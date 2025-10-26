package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"crud-backend/usuarios/dto"
	"crud-backend/usuarios/entity"
	"crud-backend/usuarios/service"

	"github.com/gorilla/mux"
)

type UsuarioHandler struct {
	service service.UsuarioService
}

func NewUsuarioHandler(s service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{service: s}
}

func (h *UsuarioHandler) RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/usuarios").Subrouter()
	s.HandleFunc("", h.GetAll).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", h.GetByID).Methods("GET")
	s.HandleFunc("", h.Create).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", h.Update).Methods("PUT")
	s.HandleFunc("/{id:[0-9]+}", h.Delete).Methods("DELETE")
}

func (h *UsuarioHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *UsuarioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	data, err := h.service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *UsuarioHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // ❌ Rechaza campos no definidos en el DTO

	var req dto.UsuarioRequest
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "❌ JSON inválido o contiene campos no permitidos", http.StatusBadRequest)
		return
	}

	// 🔍 Validar campos obligatorios (ninguno puede venir vacío)
	if req.Nombre == "" || req.Correo == "" || req.Password == "" {
		http.Error(w, "❌ Todos los campos (nombre, correo, password) son obligatorios", http.StatusBadRequest)
		return
	}

	// 🧠 Validación extra: formato del correo
	if !isValidEmail(req.Correo) {
		http.Error(w, "❌ El formato del correo electrónico no es válido", http.StatusBadRequest)
		return
	}

	u := entity.Usuario{
		Nombre:   req.Nombre,
		Correo:   req.Correo,
		Password: req.Password,
	}

	data, err := h.service.Save(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// ✅ Validador simple de correos
func isValidEmail(email string) bool {
	re := `^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(re, email)
	return matched
}

func (h *UsuarioHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.UsuarioRequest
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "JSON inválido o contiene campos no permitidos", http.StatusBadRequest)
		return
	}

	u := entity.Usuario{
		Nombre:   req.Nombre,
		Correo:   req.Correo,
		Password: req.Password,
	}
	u.ID = uint(id)

	data, err := h.service.Update(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (h *UsuarioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	// Borrado duro (definitivo)
	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
