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

	request.HandleFunc("/health", healthCheck.HealthCheckHandler).Methods("GET")
	request.HandleFunc("/getallconfigs", mongoServer.GetClientConfigAll).Methods("GET")

	// example http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc("/getconfig", mongoServer.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite).Methods("GET")

	// example http://localhost:8085/insertnew
	request.HandleFunc("/insertnew", mongoServer.InsertNewConfig).Methods("POST")

	// example http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc("/delete", mongoServer.DeleteRecordUsingID).Methods("DELETE")

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
