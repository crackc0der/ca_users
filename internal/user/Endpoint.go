package user

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type service interface {
	AddUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

type Endpoint struct {
	service service
	log     *slog.Logger
}

func (e *Endpoint) GetUser(writer http.ResponseWriter, request *http.Request) {
	var userID int
	if err := json.NewDecoder(request.Body).Decode(&userID); err != nil {
		e.log.Error(err.Error())
	}

	user, err := e.service.GetUser(request.Context(), userID)
	if err != nil {
		e.log.Error(err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&user); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) AddUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		e.log.Error(err.Error())
	}

	addedUser, err := e.service.AddUser(request.Context(), &user)
	if err != nil {
		e.log.Error(err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&addedUser); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		e.log.Error(err.Error())
	}

	updatedUser, err := e.service.UpdateUser(request.Context(), &user)
	if err != nil {
		e.log.Error(err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&updatedUser); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var userID int
	if err := json.NewDecoder(request.Body).Decode(&userID); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) GetAllUsers(writer http.ResponseWriter, _ *http.Request) {
	var users []User
	if err := json.NewEncoder(writer).Encode(&users); err != nil {
		e.log.Error(err.Error())
	}
}
