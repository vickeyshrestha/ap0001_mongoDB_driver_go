package mongoAdapter

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ap0001_mongoDB_driver_go/internal/initialConfig"
)

var mongoDbURL = initialConfig.GetMongoDBEndpoint() + ":" + initialConfig.GetMongoDBPort()

func MongoAdapterTest() {
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
