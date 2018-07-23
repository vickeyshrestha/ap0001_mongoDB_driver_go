package mongoAdapter

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ap0001_mongoDB_driver_go/internal/initialConfig"
	"net/http"
	"encoding/json"
)

/*
func MongoAdapterTest() {
	var mongoDbURL = initialConfig.GetMongoDBEndpoint() + ":" + initialConfig.GetMongoDBPort()
	session, err := mgo.Dial(mongoDbURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(Person{"Ale", "+55 53 8116 9639"},
		Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
*/

type Server struct {
	session *mgo.Session
}

func NewServer() (*Server, error) {
	var mongoDbURL = initialConfig.GetMongoDBEndpoint() + ":" + initialConfig.GetMongoDBPort()
	session, err := mgo.Dial(mongoDbURL)
	if err != nil {
		return nil, err
	}
	return &Server{session: session}, nil
}

func (s *Server) Close() {
	s.session.Close()
}

func (s *Server) GetClientConfigAll(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) GetClientConfigBasedOnAppNameAndBinaryVersionAndSite(w http.ResponseWriter, r *http.Request) {
	applicationName := r.URL.Query().Get("app")
	binaryVersion := r.URL.Query().Get("bin")
	site := r.URL.Query().Get("site")

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
	responseByte, _ := json.Marshal(clientConfig)

	if len(clientConfig) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		responseByte, _ = json.Marshal(ErrorJson{
			Error: "Cannot find the config in data store",
		})
	}
	w.Write(responseByte)
}
