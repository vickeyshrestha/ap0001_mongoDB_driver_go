package main

import (
	"ap0001_mongoDB_driver_go/internal/adapter"
	"ap0001_mongoDB_driver_go/internal/initialConfig"
)

func main() {
	initialConfig.LoadConfiguration()
	adapter.MainProcess()
}