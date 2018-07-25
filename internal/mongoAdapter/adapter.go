package mongoAdapter

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ap0001_mongoDB_driver_go/internal/initialConfig"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Server struct {
	session *mgo.Session
}

func NewServer() (*Server, error) {
	var mongoDbURL = *initialConfig.GetMongoHostAndPort()
	session, err := mgo.Dial(mongoDbURL)
	if err != nil {
		return nil, err
	}
	return &Server{session: session}, nil
}

func (s *Server) Close() {
	s.session.Close()
}

/*
	Inserts new record (configuration) into MongoDB
 */
func (s *Server) InsertNewConfig(w http.ResponseWriter, r *http.Request) {
	session := s.session.Copy()
	defer session.Close()

	var clientConfig bson.M // Since we don't know the exact structure of JSON, we will use a map instead of struct
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &clientConfig)
	_, okAppName := clientConfig["applicationName"]
	_, okBinVer := clientConfig["binaryVersion"]
	_, okSite := clientConfig["site"]
	if !(okAppName && okBinVer && okSite) {
		w.Write([]byte("Missing field. The fields applicationName, binaryVersion and site are mandatory"))
	} else {
		collection := session.DB(initialConfig.GetMongoConfigurationDatabase()).C(initialConfig.GetMongoConfigurationDbCollectionName())
		if err := collection.Insert(clientConfig); err != nil {
			panic(err)
		}
		w.Write([]byte("Successfully Inserted config"))
	}
}

/*
	Returns all records as JSON from collection
*/
func (s *Server) GetClientConfigAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := s.session.Copy()
	defer session.Close()

	//clientConfig := []ClientConfig{}
	var clientConfig []bson.M // Since we don't know the exact structure of JSON, we will use a map instead of struct
	collection := session.DB(initialConfig.GetMongoConfigurationDatabase()).C(initialConfig.GetMongoConfigurationDbCollectionName())
	err := collection.Find(bson.M{}).All(&clientConfig)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseByte, _ := json.Marshal(clientConfig)
	w.Write(responseByte)
}

/*
	Returns collection as JSON based of Application Name, Binary Version and Site
*/
func (s *Server) GetClientConfigBasedOnAppNameAndBinaryVersionAndSite(w http.ResponseWriter, r *http.Request) {
	var responseByte []byte
	applicationName := r.URL.Query().Get("app")
	binaryVersion := r.URL.Query().Get("bin")
	site := r.URL.Query().Get("site")
	if applicationName == "" || binaryVersion == "" || site == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseByte, _ = json.Marshal(ErrorJson{
			Error: "One or the more of the mandatory parameters are missing. Mandatory parameters - app, bin & site",
		})
	} else {
		session := s.session.Copy()
		defer session.Close()

		var clientConfig []bson.M
		collection := session.DB(initialConfig.GetMongoConfigurationDatabase()).C(initialConfig.GetMongoConfigurationDbCollectionName())
		err := collection.Find(bson.M{
			"applicationName": applicationName,
			"binaryVersion":   binaryVersion,
			"site":            site,
		}).All(&clientConfig)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		responseByte, _ = json.Marshal(clientConfig)

		if len(clientConfig) <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			responseByte, _ = json.Marshal(ErrorJson{
				Error: "Cannot find the config in data store",
			})
		}
	}
	w.Write(responseByte)
}

/*
	Delete record using ID
 */
func (s *Server) DeleteRecordUsingID(w http.ResponseWriter, r *http.Request) {
	var responseByte []byte
	appName := r.URL.Query().Get("app")
	binaryVersion := r.URL.Query().Get("bin")
	site := r.URL.Query().Get("site")

	if appName == "" || binaryVersion == "" || site == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseByte, _ = json.Marshal(ErrorJson{
			Error: "One or the more of the mandatory parameters are missing. Mandatory parameters - app, bin & site",
		})
	} else {
		session := s.session.Copy()
		defer session.Close()
		collection := session.DB(initialConfig.GetMongoConfigurationDatabase()).C(initialConfig.GetMongoConfigurationDbCollectionName())

		//err := collection.Remove(bson.M{"_id": id})
		err := collection.Remove(bson.M{"applicationName": appName,
			"binaryVersion": binaryVersion,
			"site": site,
		})
		if err != nil {
			msg:= "Error while removing record with params " + appName+", "+binaryVersion+" and "+site+ " | Message: "+err.Error()
			w.Write([]byte(msg))
			log.Printf(msg)
		} else {
			msg:= "Record with param " + appName+", "+binaryVersion+" and "+site+ " removed"
			w.Write([]byte(msg))
			log.Printf(msg)
		}
	}
	w.Write(responseByte)
}
