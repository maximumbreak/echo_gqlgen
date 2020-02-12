package databases

import (
	"gopkg.in/mgo.v2"
	"log"
)

func GetMGO() *mgo.Session {

	url := "localhost:27017"
	session, err := mgo.Dial(url)

	if err != nil {
		log.Print("error connection db : ", err)
		panic(err)
	} else {
		return session
	}
}
