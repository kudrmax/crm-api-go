package contacts_update

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"my/crm-golang/internal/models/contact"
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
	contactUpdateData := &contact.ContactUpdateData{}
	if err := json.NewDecoder(r.Body).Decode(contactUpdateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.Update(name, contactUpdateData)
	if errors.Is(err, my_errors.NameAlreadyUsedErr) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
