package mongoAdapter

import "gopkg.in/mgo.v2/bson"

type ClientConfig struct {
	Id              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Seqno           int           `json:"Seqno"`
	ApplicationName string        `json:"ApplicationName" bson:"applicationName,omitempty"`
	Site            string        `json:"Site" bson:"site,omitempty"`
	BinaryVersion   string        `json:"BinaryVersion"`
	ServingPort     int           `json:"ServingPort"`
}

type ErrorJson struct {
	Error string `json:"ERROR"`
}
