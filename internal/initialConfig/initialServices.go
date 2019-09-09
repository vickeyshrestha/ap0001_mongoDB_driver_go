package initialConfig

import (
	"net/http"
	"strings"
	"time"
)

// ---- BEGIN part of config file ------------

func (c configFile) GetApplicationSite() string {
	return strings.Join(c.Site, "")
}

func (c configFile) GetApplicationBinary() string {
	return strings.Join(c.BinaryVersion, "")
}

func (c configFile) GetHttpClient() http.Client {
	var httpConnectionTimeout = int32(c.HTTPConnectionTimeout)
	var client = http.Client{
		Timeout: time.Duration(httpConnectionTimeout) * time.Second,
	}
	return client
}

func (c configFile) GetMongoConfigurationDatabase() string {
	return strings.Join(c.MongoConfigurationDatabase, "")
}

func (c configFile) GetMongoConfigurationDbCollectionName() string {
	return strings.Join(c.MongoConfigurationDbCollectionName, "")
}

//-----------END part of config file ------------

func (c configFile) GetAppStartupTime() time.Time {
	return appStartUpTime
}

func (c configFile) GetMongoHostAndPort() *string {
	return mongoDbHostAndPort
}

func (c configFile) GetSslKey() *string {
	return sslKey
}

func (c configFile) GetSslCert() *string {
	return sslCert
}

func (c configFile) GetSSLMode() *string {
	return devMode
}
