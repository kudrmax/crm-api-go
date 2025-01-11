package contact_logs_get_one

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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
	logIdStr := chi.URLParam(r, "log_id")
	logId, err := strconv.Atoi(logIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logModels, err := h.contactLogService.GetByLogId(logId)

	if errors.Is(err, my_errors.ContactIdNotFoundErr) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(logModels)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
