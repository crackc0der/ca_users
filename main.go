package main

import (
	"ca/internal/user"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	//nolint
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	dsn := "postgres://postgres:example@localhost:5432/users?sslmode=disable"
	timeout := 10
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	repository, err := user.NewRepository(dsn)

	if err != nil {
		log.Fatal(err)
	}

	service := user.NewService(repository)

	endpoint := user.NewEndpoint(service, logger)

	router.HandleFunc("/", endpoint.GetAllUsers)
	router.HandleFunc("/add", endpoint.AddUser)
	router.HandleFunc("/update/{id:[0-9]+}", endpoint.UpdateUser)
	router.HandleFunc("/get/{id:[0-9]+}", endpoint.GetUser)

	srv := http.Server{
		Addr:              ":80",
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
