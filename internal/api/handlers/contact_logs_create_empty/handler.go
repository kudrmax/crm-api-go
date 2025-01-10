package contact_logs_create_empty

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"my/crm-golang/internal/models/contact_log"
	"my/crm-golang/internal/my_errors"
)

type Handler struct {
	contactLogService ContactLogService
	contactService    ContactService
}

func New(contactLogService ContactLogService, contactService ContactService) *Handler {
	return &Handler{
		contactLogService: contactLogService,
		contactService:    contactService,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	contactId, err := h.contactService.GetIdByName(name)
	if errors.Is(err, my_errors.ContactNotFoundErr) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model := new(contact_log.ContactLog)
	model.ContactId = contactId

	err = h.contactLogService.Create(model)

	if errors.Is(err, my_errors.ContactIdNotFoundErr) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
