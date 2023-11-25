package user

import (
	"context"
	"encoding/json"
	"net/http"
)

type service interface {
	AddUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
}

func NewEndpoint(service *Service) *Endpoint {
	return &Endpoint{service: service}
}

type Endpoint struct {
	service service
}

func (e *Endpoint) GetUser(writer http.ResponseWriter, request *http.Request) {
	var userID int
	if err := json.NewDecoder(request.Body).Decode(&userID); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	user, err := e.service.GetUser(request.Context(), userID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewEncoder(writer).Encode(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (e *Endpoint) AddUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	addedUser, err := e.service.AddUser(request.Context(), &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewEncoder(writer).Encode(&addedUser); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (e *Endpoint) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	updatedUser, err := e.service.UpdateUser(request.Context(), &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewEncoder(writer).Encode(&updatedUser); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (e *Endpoint) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var userID int
	if err := json.NewDecoder(request.Body).Decode(&userID); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (e *Endpoint) GetAllUsers(writer http.ResponseWriter, _ *http.Request) {
	var users []User
	if err := json.NewEncoder(writer).Encode(&users); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}
