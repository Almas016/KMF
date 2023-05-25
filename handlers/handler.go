package handlers

import (
	"KMF/entities"
	"KMF/usecases"

	"encoding/json"
	"net/http"
)

type Handler struct {
	useCase *usecases.UseCase
}

func NewHandler(useCase *usecases.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request entities.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	response := h.useCase.Process(&request)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
