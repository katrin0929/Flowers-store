package handler

import (
	"Flowers-store/internal/model"
	"Flowers-store/internal/service"
	"encoding/json"
	"net/http"
)

type RegistrationHandler struct {
	regSvc *service.RegistrationService
}

func NewRegistrationHandler(regSvc *service.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		regSvc: regSvc,
	}
}

func (rh *RegistrationHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = rh.regSvc.Register(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registered successfully"})
}
