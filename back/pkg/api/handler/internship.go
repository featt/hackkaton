package handler

import (
	"backend/pkg/model"
	"backend/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type InternshipHandler struct {
	service *service.InternshipService
}

func NewInternshipHandler(service *service.InternshipService) *InternshipHandler {
	return &InternshipHandler{service: service}
}

func (h *InternshipHandler) CreateInternship(w http.ResponseWriter, r *http.Request) {
	var internship model.Internship
	err := json.NewDecoder(r.Body).Decode(&internship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	internshipID, err := h.service.CreateInternship(r.Context(), internship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.FormatInt(internshipID, 10)))
}

func (h *InternshipHandler) GetInternship(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	internship, err := h.service.GetInternshipByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(internship)
}

func (h *InternshipHandler) UpdateInternship(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var internship model.Internship
	err = json.NewDecoder(r.Body).Decode(&internship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UpdateInternship(r.Context(), id, internship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *InternshipHandler) DownloadExcelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteInternship(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *InternshipHandler) SaveExcelHandler(w http.ResponseWriter, r *http.Request) {
	internships, err := h.service.GetAllInternships(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(internships)
}
