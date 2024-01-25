package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Angstreminus/Effective-mobile-test-task/internal/apperrors"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/dto"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/service"
	uuid "github.com/satori/go.uuid"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(serv *service.UserService) *UserHandler {
	return &UserHandler{
		Service: serv,
	}
}

func (uh *UserHandler) GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	cursor := req.URL.Query().Get("cursor")
	limitStr := req.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	filters := make(map[string]string)

	for k, v := range req.URL.Query() {
		if k != "cursor" && k != "limit" && len(v) != 0 {
			filters[k] = v[0]
		}
	}

	users, nextCursor, err := uh.Service.GetAllUsers(cursor, limit, filters)
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.MatchError(err))
	}
	data := map[string]interface{}{
		"users":       users,
		"next_cursor": nextCursor,
	}
	var resp dto.Response
	resp.Code = http.StatusOK
	resp.Data = data
	json.NewEncoder(w).Encode(resp)
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	splittedPath := strings.Split(req.URL.Path, "/")
	userId, err := uuid.FromString(splittedPath[len(splittedPath)-1])
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err = uh.Service.DeleteUser(userId)
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.MatchError(err))
	}
	var resp dto.Response
	resp.Code = http.StatusNotFound
	resp.Data = ""
	json.NewEncoder(w).Encode(resp)
}

func (uh *UserHandler) EditUser(w http.ResponseWriter, req *http.Request) {
	splittedPath := strings.Split(req.URL.Path, "/")
	userId, err := uuid.FromString(splittedPath[len(splittedPath)-1])
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var toEdit dto.UserEditRequest

	if err := json.NewDecoder(req.Body).Decode(&toEdit); err != nil {
		json.NewEncoder(w).Encode(apperrors.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer req.Body.Close()

	user, err := uh.Service.EditUser(userId, toEdit)
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.MatchError(err))
	}
	var resp dto.Response
	resp.Code = http.StatusOK
	resp.Data = user
	json.NewEncoder(w).Encode(resp)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	var toCreate dto.UserRequest
	err := json.NewDecoder(req.Body).Decode(&toCreate)
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	user, err := uh.Service.CreateUser(toCreate)
	if err != nil {
		json.NewEncoder(w).Encode(apperrors.MatchError(err))
	}
	var resp dto.Response
	resp.Code = http.StatusCreated
	resp.Data = user
	json.NewEncoder(w).Encode(resp)
}
