package ap0001_mongo_engine

const (
	ApplicationName = "MongodbDriver"

	HealthCheck               = "/health"
	GetAllConfigsFromDatabase = "/getallconfigs"
	GetSingleConfig           = "/getconfig"
	InsertConfig              = "/insertnew"
	DeleteConfig              = "/delete"
)
