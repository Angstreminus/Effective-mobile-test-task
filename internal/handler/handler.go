package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Angstreminus/Effective-mobile-test-task/internal/dto"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/service"
	uuid "github.com/satori/go.uuid"
)

type UserHandler struct {
	Service *service.UserService
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
		// json.NewEncoder(err)
	}
	response := map[string]interface{}{
		"users":       users,
		"next_cursor": nextCursor,
	}
	json.NewEncoder(w).Encode(response)
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	splittedPath := strings.Split(req.URL.Path, "/")
	userId, err := uuid.FromString(splittedPath[len(splittedPath)-1])
	if err != nil {
		// write err
	}

	err = uh.Service.DeleteUser(userId)
	if err != nil {
	}

	//json.NewEncoder
}

func (uh *UserHandler) EditUser(w http.ResponseWriter, req *http.Request) {
	splittedPath := strings.Split(req.URL.Path, "/")
	userId, err := uuid.FromString(splittedPath[len(splittedPath)-1])
	if err != nil {
	}

	var toEdit dto.UserEditRequest

	if err := json.NewDecoder(req.Body).Decode(&toEdit); err != nil {

	}
	defer req.Body.Close()

	user, err := uh.Service.EditUser(userId, toEdit)
	if err != nil {
		// json.Error
	}
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	var toCreate dto.UserRequest
	err := json.NewDecoder(req.Body).Decode(&toCreate)
	if err != nil {

	}

	user, err := uh.Service.CreateUser(toCreate)

	json.NewEncoder(w).Encode(user)
}
