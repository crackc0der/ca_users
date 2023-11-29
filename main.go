package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	//nolint
	"github.com/crackc0der/users/config"
	//nolint
	"github.com/crackc0der/users/internal/user"
	//nolint
	"github.com/gorilla/mux"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	dsn := "postgres://" + config.DataBase.DBUser + ":" + config.DataBase.DBPassword + "@" + config.DataBase.DBHost +
		":" + config.DataBase.DBPort + "/" + config.DataBase.DBName + "?sslmode=disable"
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
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
