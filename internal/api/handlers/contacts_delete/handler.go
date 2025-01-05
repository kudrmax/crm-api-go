package contacts_delete

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

	err := h.service.DeleteByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
