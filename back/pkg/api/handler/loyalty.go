package handler

import (
	"backend/pkg/model"
	"backend/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type LoyaltyHandler struct {
	service *service.LoyaltyService
}

func NewLoyaltyHandler(service *service.LoyaltyService) *LoyaltyHandler {
	return &LoyaltyHandler{
		service: service,
	}
}

func (h *LoyaltyHandler) GetUserLoyalty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loyaltyPoints, err := h.service.GetUserLoyalty(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"loyalty_points": loyaltyPoints})
}
func (h *LoyaltyHandler) UpdateUserLoyalty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var loyaltyReq model.UpdateLoyaltyRequest
	err = json.NewDecoder(r.Body).Decode(&loyaltyReq)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UpdateUserLoyalty(r.Context(), id, loyaltyReq.LoyaltyPoints)
	if err != nil {
		log.Printf("Error updating user loyalty: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
