package main

import (
	"ca/internal/user"
	"log"
	"net/http"
	"time"

	//nolint
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	dsn := "postgres"
	timeout := 10

	repository, err := user.NewRepository(dsn)

	if err != nil {
		log.Fatal(err)
	}

	service := user.NewService(repository)

	endpoint := user.NewEndpoint(service)

	router.HandleFunc("/", endpoint.GetAllUsers)

	srv := http.Server{
		Addr:              ":80",
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
