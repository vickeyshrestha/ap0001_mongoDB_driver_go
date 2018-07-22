package main

import (
	"ap0001_mongoDB_driver_go/internal/initialConfig"
	"ap0001_mongoDB_driver_go/internal/applicationDriver"
)

func main() {
	initialConfig.LoadConfiguration()
	applicationDriver.MainProcess()
	//adapter.MongoAdapterTest() // Later will be placed by Main Process
}