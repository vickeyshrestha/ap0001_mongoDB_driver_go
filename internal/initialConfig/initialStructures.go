package initialConfig

type configFileStruct struct {
	Site                  []string `json:"site"`
	BinaryVersion         []string `json:"binary_version"`
	MongoDbEndpoint       []string `json:"mongo_db_endpoint"`
	MongoDbPort           []string `json:"mongo_db_port"`
	HTTPConnectionTimeout int      `json:"http_connection_timeout"`
}
