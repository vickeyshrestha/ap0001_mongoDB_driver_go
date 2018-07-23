package initialConfig

import (
	"strings"
	"net/http"
	"time"
)

// ---- BEGIN part of config file ------------
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

func GetMongoConfigurationDatabase() string {
	var mongoConfigurationDatabase = strings.Join(configFromJsonFile.MongoConfigurationDatabase,"")
	return mongoConfigurationDatabase
}

func GetMongoConfigurationDbCollectionName() string {
	var mongoConfigurationDbCollectionName = strings.Join(configFromJsonFile.MongoConfigurationDbCollectionName,"")
	return mongoConfigurationDbCollectionName
}

//-----------END part of config file ------------

func GetAppStartupTime() time.Time {
	return appStartUpTime
}

