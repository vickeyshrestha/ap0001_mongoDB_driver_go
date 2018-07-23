package healthCheck

import (
	"net/http"
	"encoding/json"
)

func HealthCheckHandler (writer http.ResponseWriter, request *http.Request){
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	responseByte, _ := json.Marshal(healthEndpoint{
		Application: "MongoDB Driver",
		HealthStatus: "200",
		Message: "Up and running",
	})
	writer.Write(responseByte)
}
