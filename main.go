package main

import (
	"ca/internal/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	dsn := "postgres"
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
		ReadHeaderTimeout: time.Second * 10,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
