package applicationDriver

import (
	"github.com/gorilla/mux"
	"ap0001_mongoDB_driver_go/internal/healthCheck"
	"gopkg.in/tylerb/graceful.v1"
	"time"
	"net/http"
	"log"
	"ap0001_mongoDB_driver_go/internal/mongoAdapter"
)

func MainProcess() {
	request := mux.NewRouter().StrictSlash(false)

	mongoServer, err := mongoAdapter.NewServer()
	if err != nil {
		panic(err)
	}
	defer mongoServer.Close()

	request.HandleFunc("/health", healthCheck.HealthCheckHandler)
	request.HandleFunc("/getallconfigs", mongoServer.GetClientConfigAll)
	request.HandleFunc("/getconfig", mongoServer.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite)

	server:= &graceful.Server{
		Timeout: 30 * time.Second,
		Server : &http.Server{
			Addr: ":8085",
			Handler: request,
		},
	}

	log.Printf("Application started successfully. Serving port 8085")

	server.ListenAndServe()
	if err!= nil {
		log.Printf("Server failed to start | %v", err)
	}

	log.Printf("Application stopped gracefully")
}
