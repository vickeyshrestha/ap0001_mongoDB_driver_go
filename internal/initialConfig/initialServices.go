package initialConfig

import (
	"net/http"
	"strings"
	"time"
)

// ---- BEGIN part of config file ------------

func (c configFileStruct) GetApplicationSite() string {
	return strings.Join(c.Site, "")
}

func (c configFileStruct) GetApplicationBinary() string {
	return strings.Join(c.BinaryVersion, "")
}

func (c configFileStruct) GetHttpClient() http.Client {
	var httpConnectionTimeout = int32(c.HTTPConnectionTimeout)
	var client = http.Client{
		Timeout: time.Duration(httpConnectionTimeout) * time.Second,
	}
	return client
}

func (c configFileStruct) GetMongoConfigurationDatabase() string {
	return strings.Join(c.MongoConfigurationDatabase, "")
}

func (c configFileStruct) GetMongoConfigurationDbCollectionName() string {
	return strings.Join(c.MongoConfigurationDbCollectionName, "")
}

//-----------END part of config file ------------

func (c configFileStruct) GetAppStartupTime() time.Time {
	return appStartUpTime
}

func (c configFileStruct) GetMongoHostAndPort() *string {
	return mongoDbHostAndPort
}

func (c configFileStruct) GetSslKey() *string {
	return sslKey
}

func (c configFileStruct) GetSslCert() *string {
	return sslCert
}

func (c configFileStruct) GetSSLMode() *string {
	return devMode
}
