package ap0001_mongo_engine

import (
	"net/http"
	"time"
)

type InitialConfig interface {
	GetApplicationSite() string
	GetApplicationBinary() string
	GetHttpClient() http.Client
	GetMongoConfigurationDatabase() string
	GetMongoConfigurationDbCollectionName() string
	GetAppStartupTime() time.Time
	GetMongoHostAndPort() *string
	GetSslKey() *string
	GetSslCert() *string
	GetSSLMode() *string
}

type HealthHandler interface {
	HealthCheck(http.ResponseWriter, *http.Request)
}

type MongoAdapter interface {
	InsertNewConfig(http.ResponseWriter, *http.Request)
	GetClientConfigAll(http.ResponseWriter, *http.Request)
	GetClientConfigBasedOnAppNameAndBinaryVersionAndSite(http.ResponseWriter, *http.Request)
	DeleteRecordUsingID(http.ResponseWriter, *http.Request)
	Close()
}
