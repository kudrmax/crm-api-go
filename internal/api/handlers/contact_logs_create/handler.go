package contact_logs_create

import (
	"encoding/json"
	"errors"
	"net/http"

	"my/crm-golang/internal/models/contact_log"
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
	model := new(contact_log.ContactLog)
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.Create(model)

	if errors.Is(err, my_errors.ContactIdNotFoundErr) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
