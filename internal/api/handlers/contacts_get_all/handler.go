package contacts_get_all

import (
	"encoding/json"
	"net/http"
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
	contactModels, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response := map[string]interface{}{"data": contactModel}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contactModels)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
