package main

import (
	"ap0001_mongo_engine/internal/initialConfig"
	"ap0001_mongo_engine/internal/applicationDriver"
)

func main() {
	initialConfig.LoadConfiguration()
	applicationDriver.MainProcess()
	//mongoAdapter.MongoAdapterTest() // Later will be placed by Main Process
}