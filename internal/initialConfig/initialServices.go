package initialConfig

import (
	"strings"
	"net/http"
	"time"
)

func GetApplicationSite() string {
	var appSite string
	appSite = strings.Join(configFromJsonFile.Site,"")
	return appSite
}

func GetApplicationBinary() string {
	var appBinary string
	appBinary = strings.Join(configFromJsonFile.BinaryVersion,"")
	return appBinary
}

func GetMongoDBEndpoint() string {
	var mongoEndpoint string
	mongoEndpoint = strings.Join(configFromJsonFile.MongoDbEndpoint,"")
	return mongoEndpoint
}

func GetMongoDBPort() string {
	var mongoPort string
	mongoPort = strings.Join(configFromJsonFile.MongoDbPort,"")
	return mongoPort
}

func GetHttpClient() http.Client {
	var httpConnectionTimeout = int32(configFromJsonFile.HTTPConnectionTimeout)
	var client = http.Client{
		Timeout: time.Duration(httpConnectionTimeout) * time.Second,
	}
	return client
}

func GetAppStartupTime() time.Time {
	return appStartUpTime
}

