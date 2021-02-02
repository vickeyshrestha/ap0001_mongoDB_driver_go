package main

import (
	"ap0001_mongo_engine-DEPRECIATED/internal/applicationDriver"
	"ap0001_mongo_engine-DEPRECIATED/internal/initialConfig"
	"log"
)

func main() {
	config, err := initialConfig.NewConfiguration()
	if err != nil {
		log.Println(err)
	}
	applicationDriver.Start(config)
}
