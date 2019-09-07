package main

import (
	"ap0001_mongo_engine/internal/applicationDriver"
	"ap0001_mongo_engine/internal/initialConfig"
	"log"
)

func main() {
	config, err := initialConfig.NewConfiguration()
	if err != nil {
		log.Println(err)
	}
	applicationDriver.Start(config)
}
