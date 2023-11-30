package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/crackc0der/users/config"
	"github.com/crackc0der/users/internal/user"
	"github.com/gorilla/mux"
)

func main() {
	config, err := config.NewConfig()
	dsn := config.GetDsn()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
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
	router.HandleFunc("/update", endpoint.UpdateUser)
	router.HandleFunc("/delete/{id:[0-9]+}", endpoint.DeleteUser)
	router.HandleFunc("/get/{id:[0-9]+}", endpoint.GetUser)

	//nolint
	srv := http.Server{
		Addr:              config.Host.HostPort,
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
