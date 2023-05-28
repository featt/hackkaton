// handler/user.go
package handler

import (
	"backend/pkg/model"
	"backend/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
	sessionService *service.SessionService
}

func NewUserHandler(service *service.UserService, sessionService *service.SessionService) *UserHandler {
	return &UserHandler{
		service:        service,
		sessionService: sessionService,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var regReq model.RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.RegisterUser(r.Context(), regReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.LoginUser(r.Context(), loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UpdateUser(r.Context(), id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	var roleReq model.UpdateRoleRequest
	err := json.NewDecoder(r.Body).Decode(&roleReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UpdateUserRole(r.Context(), roleReq.UserID, string(roleReq.Role))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userIdStr, err := h.sessionService.ExtractUserIdFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
