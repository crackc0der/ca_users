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
	return &Endpoint{service: *service}
}

type Endpoint struct {
	service Service
}

func (e *Endpoint) GetUser(writer http.ResponseWriter, request *http.Request) {
	var userID int
	if err := json.NewDecoder(request.Body).Decode(&userID); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	user, err := e.service.GetUser(request.Context(), 45)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(writer).Encode(user)
}

func (e *Endpoint) AddUser(writer http.ResponseWriter, request *http.Request) {

}

func (e *Endpoint) UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func (e *Endpoint) DeleteUser(writer http.ResponseWriter, request *http.Request) {

}

func (e *Endpoint) GetAllUsers(writer http.ResponseWriter, request *http.Request) {

}
