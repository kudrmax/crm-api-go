package contacts_get_one

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"my/crm-golang/internal/my_errors"
)

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	}

	contactModel, err := h.service.GetByName(name)
	if errors.Is(err, my_errors.ContactNotFoundErr) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response := map[string]interface{}{"data": contactModel}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contactModel)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
