package healthCheck

import (
	"ap0001_mongo_engine"
	"encoding/json"
	"net/http"
	"time"
)

type Service struct {
	config ap0001_mongo_engine.InitialConfig
}

func NewHealthService(config ap0001_mongo_engine.InitialConfig) (ap0001_mongo_engine.HealthHandler, error) {
	return &Service{config: config}, nil
}

func (s *Service) HealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	responseByte, _ := json.Marshal(HealthEndpoint{
		Application:  "Mongo Engine",
		Version:      s.config.GetApplicationBinary(),
		HealthStatus: "200 OK",
		Message:      "Up and running for " + time.Since(s.config.GetAppStartupTime()).String(),
	})
	_, _ = writer.Write(responseByte)
}
