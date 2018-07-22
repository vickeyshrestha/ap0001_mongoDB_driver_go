package initialConfig

type configFileStruct struct {
	ApplicationName       string `json:"application_name"`
	Site                  string `json:"site"`
	BinaryVersion         string `json:"binary_version"`
	MongoDbEndpoint       string `json:"mongo_db_endpoint"`
	MongoDbPort           string `json:"mongo_db_port"`
	HTTPConnectionTimeout string `json:"http_connection_timeout"`
}
