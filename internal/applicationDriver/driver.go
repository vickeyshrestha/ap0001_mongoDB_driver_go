package applicationDriver

import (
	"github.com/gorilla/mux"
	"ap0001_mongo_engine/internal/healthCheck"
	"gopkg.in/tylerb/graceful.v1"
	"time"
	"net/http"
	"log"
	"ap0001_mongo_engine/internal/mongoAdapter"
	"ap0001_mongo_engine/internal/generalUtilities"
	"fmt"
	"os"
	"ap0001_mongo_engine/internal/controllers"
	"strings"
	"ap0001_mongo_engine/internal/initialConfig"
)

func Start() {
	request := mux.NewRouter().StrictSlash(false)

	mongoServer, err := mongoAdapter.NewServer()
	if err != nil {
		log.Printf("Cannot connecto to MongoDB. ERROR: %v", err)
		os.Exit(1)
	} else {
		defer mongoServer.Close()

		// example: http://localhost:8085/health
		request.HandleFunc(controllers.HEALTH_CHECK, healthCheck.HealthCheckHandler).Methods("GET")

		// example: http://localhost:8085/getallconfigs
		request.HandleFunc(controllers.GET_ALL_CONFIGS_FROM_DATABASE, mongoServer.GetClientConfigAll).Methods("GET")

		// example http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev
		request.HandleFunc(controllers.GET_SINGLE_CONFIG, mongoServer.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite).Methods("GET")

		// example http://localhost:8085/insertnew
		request.HandleFunc(controllers.INSERT_CONFIG, mongoServer.InsertNewConfig).Methods("POST")

		// example http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
		request.HandleFunc(controllers.DELETE_CONFIG, mongoServer.DeleteRecordUsingID).Methods("DELETE")

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
		if strings.EqualFold(*initialConfig.GetSSLMode(),"false"){
			log.Printf("Dev mode set to false. Starting application in ssl secured mode")
			errStartingServer := server.ListenAndServeTLS(*initialConfig.GetSslCert(), *initialConfig.GetSslKey())
			if errStartingServer != nil {
				log.Printf("Failed to start server | Error: %v", errStartingServer)
			}
		} else {
			log.Printf("Starting application in ssl non-secured mode")
			server.ListenAndServe()
		}
		log.Printf("Application stopped gracefully")
	}

}
