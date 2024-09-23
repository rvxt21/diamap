package main

import (
	"diamap/api/pkg/handlers"
	"diamap/api/pkg/services"
	"diamap/api/pkg/storage"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Diamap Project Started")
	connStr := os.Getenv("POSTGRES_CONN_STR")
	if connStr == "" {
		log.Fatal().Msgf("Environment variable POSTGRES_CONN_STR is required")
	}
	storage, err := storage.New(connStr)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
	defer storage.DB.Close()

	service := services.Service{Storage: storage}
	handler := handlers.Handler{Service: service}

	router := mux.NewRouter()
	handler.RegisterRoutes(router)
	fmt.Println("Starting server at :8080")
	errServ := http.ListenAndServe(":8080", router)
	if errServ != nil {
		fmt.Println("Error happened, %v", errServ.Error)
		return
	}
}
