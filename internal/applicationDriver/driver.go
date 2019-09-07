package applicationDriver

import (
	"ap0001_mongo_engine"
	"ap0001_mongo_engine/internal/generalUtilities"
	"ap0001_mongo_engine/internal/healthCheck"
	"ap0001_mongo_engine/internal/mongoAdapter"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/tylerb/graceful.v1"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Start(config ap0001_mongo_engine.InitialConfig) {
	request := mux.NewRouter().StrictSlash(false)

	mongoServer, err := mongoAdapter.NewServer(config)
	if err != nil {
		log.Printf("Cannot connecto to MongoDB. ERROR: %v", err)
		os.Exit(1)
	} else {
		defer mongoServer.Close()

		healthServer, err := healthCheck.NewHealthService(config)
		if err != nil {
			panic(err)
		}
		// example: http://localhost:8085/health
		request.HandleFunc(ap0001_mongo_engine.HealthCheck, healthServer.HealthCheckHandler).Methods("GET")

		// example: http://localhost:8085/getallconfigs
		request.HandleFunc(ap0001_mongo_engine.GetAllConfigsFromDatabase, mongoServer.GetClientConfigAll).Methods("GET")

		// example http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev
		request.HandleFunc(ap0001_mongo_engine.GetSingleConfig, mongoServer.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite).Methods("GET")

		// example http://localhost:8085/insertnew
		request.HandleFunc(ap0001_mongo_engine.InsertConfig, mongoServer.InsertNewConfig).Methods("POST")

		// example http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
		request.HandleFunc(ap0001_mongo_engine.DeleteConfig, mongoServer.DeleteRecordUsingID).Methods("DELETE")

		server := &graceful.Server{
			Timeout: 30 * time.Second,
			Server: &http.Server{
				Addr:    ":8085",
				Handler: request,
			},
		}

		ip, err := generalUtilities.ExternalIP()
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("Application started successfully. Running in ip %v & serving port 8085", ip)
		if strings.EqualFold(*config.GetSSLMode(), "false") {
			log.Printf("Dev mode set to false. Starting application in ssl secured mode")
			errStartingServer := server.ListenAndServeTLS(*config.GetSslCert(), *config.GetSslKey())
			if errStartingServer != nil {
				log.Printf("Failed to start server | Error: %v", errStartingServer)
			}
		} else {
			log.Printf("Starting application in ssl non-secured mode")
			err = server.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}
		log.Printf("Application stopped gracefully")
	}

}
