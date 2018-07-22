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

type ClientConfig struct {
	Id              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Seqno           int           `json:"Seqno"`
	ApplicationName string        `json:"ApplicationName" bson:"applicationName,omitempty"`
	Site            string        `json:"Site" bson:"site,omitempty"`
	BinaryVersion   string        `json:"BinaryVersion"`
	ServingPort     int           `json:"ServingPort"`
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

func (s *Server) GetClientConfig(w http.ResponseWriter, r *http.Request) {
	session := s.session.Copy()
	defer session.Close()

	clientConfig := []ClientConfig{}
	collection := session.DB("config").C("vic_application")
	err := collection.Find(bson.M{}).All(&clientConfig)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	j, _ := json.Marshal(clientConfig)
	w.Write(j)
}
