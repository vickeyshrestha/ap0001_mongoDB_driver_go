package healthCheck

import (
	"net/http"
	"encoding/json"
	"ap0001_mongo_engine/internal/initialConfig"
	"time"
)

func HealthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	responseByte, _ := json.Marshal(HealthEndpoint{
		Application:  "Mongo Engine",
		Version:      initialConfig.GetApplicationBinary(),
		HealthStatus: "200 OK",
		Message:      "Up and running for " + time.Since(initialConfig.GetAppStartupTime()).String(),
	})
	writer.Write(responseByte)
}
