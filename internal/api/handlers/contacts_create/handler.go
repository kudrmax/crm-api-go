package contacts_create

import (
	"encoding/json"
	"errors"
	"net/http"

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
	contactModel := &contact.Contact{}
	if err := json.NewDecoder(r.Body).Decode(contactModel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.Create(contactModel)
	if errors.Is(err, my_errors.NameAlreadyUsedErr) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
