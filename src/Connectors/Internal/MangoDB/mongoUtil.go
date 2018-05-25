package MangoDB

import (
	"gopkg.in/mgo.v2"
	"log"
)

func Mongo_getSession() {
	log.Println("getting session...")
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
}
