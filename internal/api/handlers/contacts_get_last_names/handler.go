package contacts_get_last_names

import (
	"encoding/json"
	"log"
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
	names, err := h.service.GetLastContactsNames(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//response := map[string]interface{}{"data": contactModel}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(names)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
