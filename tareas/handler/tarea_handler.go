package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"crud-backend/tareas/dto"
	"crud-backend/tareas/entity"
	"crud-backend/tareas/service"

	"github.com/gorilla/mux"
)

type TareaHandler struct {
	service service.TareaService
}

func NewTareaHandler(s service.TareaService) *TareaHandler {
	return &TareaHandler{service: s}
}

func (h *TareaHandler) RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/tareas").Subrouter()
	s.HandleFunc("", h.GetAll).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", h.GetByID).Methods("GET")
	s.HandleFunc("", h.Create).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", h.Update).Methods("PUT")
	s.HandleFunc("/{id:[0-9]+}", h.Delete).Methods("DELETE")
}

// ========================== CRUD ==========================

func (h *TareaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *TareaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	data, err := h.service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *TareaHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // ❌ Rechaza campos no definidos

	var req dto.TareaRequest
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "❌ JSON inválido o contiene campos no permitidos", http.StatusBadRequest)
		return
	}

	// 🔍 Validar campos obligatorios
	if req.Titulo == "" || req.Descripcion == "" {
		http.Error(w, "❌ Los campos 'titulo' y 'descripcion' son obligatorios", http.StatusBadRequest)
		return
	}

	// 🧠 Validar que el título tenga al menos 3 caracteres
	if len(req.Titulo) < 3 {
		http.Error(w, "❌ El título debe tener al menos 3 caracteres", http.StatusBadRequest)
		return
	}

	// 🧠 Validar que la descripción no sea solo espacios
	if isOnlySpaces(req.Descripcion) {
		http.Error(w, "❌ La descripción no puede estar vacía o solo contener espacios", http.StatusBadRequest)
		return
	}

	t := entity.Tarea{
		Titulo:      req.Titulo,
		Descripcion: req.Descripcion,
		Completada:  req.Completada,
	}

	data, err := h.service.Save(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *TareaHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.TareaRequest
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "❌ JSON inválido o contiene campos no permitidos", http.StatusBadRequest)
		return
	}

	if req.Titulo == "" || req.Descripcion == "" {
		http.Error(w, "❌ Los campos 'titulo' y 'descripcion' son obligatorios", http.StatusBadRequest)
		return
	}

	t := entity.Tarea{
		Titulo:      req.Titulo,
		Descripcion: req.Descripcion,
		Completada:  req.Completada,
	}
	t.ID = uint(id)

	data, err := h.service.Update(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *TareaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ========================== Helpers ==========================

// Verifica si una cadena contiene solo espacios
func isOnlySpaces(s string) bool {
	re := regexp.MustCompile(`^\s*$`)
	return re.MatchString(s)
}
