package initialConfig

type configFile struct {
	Site                               []string `json:"site"`
	BinaryVersion                      []string `json:"binary_version"`
	HTTPConnectionTimeout              int      `json:"http_connection_timeout"`
	MongoConfigurationDatabase         []string `json:"mongoConfigurationDatabase"`
	MongoConfigurationDbCollectionName []string `json:"mongoConfigurationDbCollectionName"`
}
