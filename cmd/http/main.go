package main

import (
	"fmt"
	"github.com/converge/k1-create-app-go-template/internal/handler/health"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	healthHandler := health.NewHealth()
	r.HandleFunc("/app/health", healthHandler.Health).Methods(http.MethodGet)

	port := ":7001"
	fmt.Printf("API listening at port %s\n", port[1:])
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

}
