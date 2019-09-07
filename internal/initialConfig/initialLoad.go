package initialConfig

import (
	"ap0001_mongo_engine"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

var (
	appStartUpTime = time.Now()

	//var for config file. We will read as environment variable. eg. configFile=C:\Projects-Golang\src\ap0001_mongoDB_driver_go\resources\config.json
	configJsonFile, _  = os.Open(os.Getenv("configFile"))
	mongoDbHostAndPort = flag.String("mongoHostAndPort", os.Getenv("mongoHostAndPort"), "Path for mongo db endpoint")
	sslKey             = flag.String("sslKey", os.Getenv("sslKey"), "Path for sslKey")
	sslCert            = flag.String("sslCert", os.Getenv("sslCert"), "Path for sslCert")
	devMode            = flag.String("devmode", os.Getenv("devmode"), "Check for dev mode")
)

func NewConfiguration() (ap0001_mongo_engine.InitialConfig, error) {
	log.Printf("%v | INFO: %v | Reading config file from application resources.....", time.Now().Format(time.RFC1123), ap0001_mongo_engine.ApplicationName)
	configFromJsonFile := configFileStruct{}
	decoderConfigFile := json.NewDecoder(configJsonFile)
	errDecode := decoderConfigFile.Decode(&configFromJsonFile)
	if errDecode != nil {
		log.Printf("%v | ERROR: %v | Failed to read the application config json file. Does the file exist or has the env var been set? | ERROR: %v", time.Now().Format(time.RFC1123), ap0001_mongo_engine.ApplicationName, errDecode)
		log.Printf("%v | ERROR: %v | Exiting application .... ", time.Now().Format(time.RFC1123), ap0001_mongo_engine.ApplicationName)
		return nil, errDecode
	}
	log.Printf("%v | INFO: %v | Successfully read config file after %v, ", time.Now().Format(time.RFC1123), ap0001_mongo_engine.ApplicationName, time.Since(appStartUpTime))
	return configFromJsonFile, nil

}
