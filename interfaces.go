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
