package handler

import (
	"backend/pkg/model"
	"backend/pkg/service"
	"encoding/json"
	"net/http"
)

type CandidateHandler struct {
	service *service.CandidateService
}

func NewCandidateHandler(service *service.CandidateService) *CandidateHandler {
	return &CandidateHandler{
		service: service,
	}
}

func (h *CandidateHandler) ProcessCandidateForm(w http.ResponseWriter, r *http.Request) {
	var candidate model.CandidateModel
	err := json.NewDecoder(r.Body).Decode(&candidate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	isValid := h.service.IsCandidateValid(&candidate)
	if isValid {
		err = h.service.SaveCandidate(r.Context(), &candidate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "Formal stage passed."})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "Formal stage failed."})
	}
}
