package healthCheck

import (
	"net/http"
	"io"
)

func HealthCheckHandler (writer http.ResponseWriter, request *http.Request){
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, `{"alive": true}`)
}

/*
func HealthCheckHandler1(url string) http.HandlerFunc {
	internalServer := func(writer http.ResponseWriter, request *http.Request) {

	}
}
*/