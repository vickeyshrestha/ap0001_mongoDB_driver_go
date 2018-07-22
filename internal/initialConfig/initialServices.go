package initialConfig

import "strings"

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

func GetHttpConnectionTimeout() string {
	var httpConnectionTimeout string
	httpConnectionTimeout = strings.Join(configFromJsonFile.HTTPConnectionTimeout,"")
	return httpConnectionTimeout
}


