package applicationDriver

import (
	"ap0001_mongo_engine-DEPRECIATED"
	"github.com/gorilla/mux"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"time"
)

type Service struct {
	mongo  ap0001_mongo_engine.MongoAdapter
	health ap0001_mongo_engine.HealthHandler
}

func NewService(mongoServer ap0001_mongo_engine.MongoAdapter, healthServer ap0001_mongo_engine.HealthHandler) *Service {
	return &Service{
		mongo:  mongoServer,
		health: healthServer,
	}
}

func (s *Service) Routes(request *mux.Router) *graceful.Server {
	// example: http://localhost:8085/health
	request.HandleFunc(ap0001_mongo_engine.HealthCheck, s.health.HealthCheck).Methods("GET")

	// example: http://localhost:8085/getallconfigs
	request.HandleFunc(ap0001_mongo_engine.GetAllConfigsFromDatabase, s.mongo.GetClientConfigAll).Methods("GET")

	// example http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc(ap0001_mongo_engine.GetSingleConfig, s.mongo.GetClientConfigBasedOnAppNameAndBinaryVersionAndSite).Methods("GET")

	// example http://localhost:8085/insertnew
	request.HandleFunc(ap0001_mongo_engine.InsertConfig, s.mongo.InsertNewConfig).Methods("POST")

	// example http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
	request.HandleFunc(ap0001_mongo_engine.DeleteConfig, s.mongo.DeleteRecordUsingID).Methods("DELETE")

	server := &graceful.Server{
		Timeout: 30 * time.Second,
		Server: &http.Server{
			Addr:    ":8085",
			Handler: request,
		},
	}
	return server
}
