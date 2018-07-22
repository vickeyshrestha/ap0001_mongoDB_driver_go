package initialConfig

import (
	"time"
	"os"
	"log"
	"encoding/json"
)

const applicationName  = "MongodbDriver"

var(
	appStartUpTime = time.Now()

	//var for config file. We will read as environment variable.
	// eg. configFile=C:\Projects-Golang\src\ap0001_mongoDB_driver_go\resources\config.json
	configJsonFile, _ = os.Open(os.Getenv("configFile"))
	configFromJsonFile = configFileStruct{}
)

func LoadConfiguration() {
	log.Printf("%v | INFO: %v | Reading config file from application resources.....", time.Now().Format(time.RFC1123), applicationName)
	decoderConfigFile := json.NewDecoder(configJsonFile)
	errDecode := decoderConfigFile.Decode(&configFromJsonFile)
	if errDecode != nil {
		log.Printf("%v | ERROR: %v | Failed to read the application config json file. Does the file exist or has the env var been set? | ERROR: %v", time.Now().Format(time.RFC1123), applicationName, errDecode)
		log.Printf("%v | ERROR: %v | Exiting application .... ",time.Now().Format(time.RFC1123), applicationName)
		os.Exit(1)
	}
}
