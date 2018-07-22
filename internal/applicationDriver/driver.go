package applicationDriver

import (
	"github.com/gorilla/mux"
	"ap0001_mongoDB_driver_go/internal/healthCheck"
	"gopkg.in/tylerb/graceful.v1"
	"time"
	"net/http"
	"log"
)

func MainProcess() {
	request := mux.NewRouter().StrictSlash(false)

	request.HandleFunc("/health", healthCheck.HealthCheckHandler)

	server:= &graceful.Server{
		Timeout: 30 * time.Second,
		Server : &http.Server{
			Addr: ":8085",
			Handler: request,
		},
	}

	log.Printf("Application started successfully. Serving port 8085")

	err:= server.ListenAndServe()
	if err!= nil {
		log.Printf("Server failed to start | %v", err)
	}

	log.Printf("Application stopped gracefully")
}
