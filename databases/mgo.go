package databases

import (
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
)

type BaseMgo struct {
	*mgo.Session
	*mgo.Database
}

func (b *BaseMgo) Init() {

	url := viper.GetString("mongodb.connection")
	dbName := viper.GetString("mongodb.dbName")
	session, err := mgo.Dial(url)
	if err != nil {
		log.Print("error connection db : ", err)
		panic(err)
	} else {
		b.Session = session
		b.Database = b.Session.DB(dbName)
	}

}

func (b *BaseMgo) Close() {
	b.Session.Close()
}

func GetMGO() *BaseMgo {
	db := BaseMgo{}
	db.Init()
	return &db
}
