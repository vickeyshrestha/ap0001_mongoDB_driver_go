package initialConfig

import (
	"ap0001_mongo_engine-DEPRECIATED"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

var (
	appStartUpTime = time.Now()

	// Setting up environment variables during application startup

	// 1. Location of config file, usually located at \src\ap0001_mongoDB_driver_go\resources\config.json
	configJsonFile, _ = os.Open(os.Getenv("configFile"))

	// 2. FQDN and the port where MongoDB is running
	mongoDbHostAndPort = flag.String("mongoHostAndPort", os.Getenv("mongoHostAndPort"), "Path for mongo db endpoint")

	// 3. If running in secure mode (i.e. devMode set to false), this is Path to the ssl private Key
	sslKey = flag.String("sslKey", os.Getenv("sslKey"), "Path for sslKey")

	// 4. If running in secure mode (i.e. devMode set to false), this is Path to the ssl private certificatee
	sslCert = flag.String("sslCert", os.Getenv("sslCert"), "Path for sslCert")

	// 5. If set it to false, this will enable the https secure server
	devMode = flag.String("devmode", os.Getenv("devmode"), "Check for dev mode")
)

func NewConfiguration() (ap0001_mongo_engine.InitialConfig, error) {
	log.Printf("%v | INFO: %v | Reading config file from application resources.....", time.Now().Format(time.RFC1123), ap0001_mongo_engine.ApplicationName)
	configFromJsonFile := configFile{}
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
