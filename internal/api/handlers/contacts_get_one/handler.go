package contacts_get_one

import (
	"encoding/json"
	"errors"
	"net/http"

	"my/crm-golang/internal/storage/postgres/contacts"
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
	name := r.URL.Path[len("/get/"):]
	if name == "" {
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	}

	contactModel, err := h.service.GetByName(name)
	if errors.Is(err, contacts.ContactNotFoundErr) {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Some error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"data": contactModel}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
