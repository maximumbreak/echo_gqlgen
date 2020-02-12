package api

import (
	"log"

	"github.com/beforesecond/gqlgen-todos/databases"
	"github.com/beforesecond/gqlgen-todos/models"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

const kindUser = "User"

// FindUser from datastore
func FindUser(username, password string) (*models.UserModel, error) {
	session := databases.GetMGO()
	defer session.Close()

	col := session.DB("test").C("users")
	user := models.UserModel{}
	query := bson.M{
		"user.username": username,
	}

	err := col.Find(query).One(&user)

	if err != nil {
		log.Print(username, " ", err)
	} else {
		log.Print("user =>", user)
	}
	if !user.ComparePassword(password) {
		// wrong password return like user not found
		return nil, nil
	}
	return &user, nil

}

// SaveUser to datastore
func SaveUser(user *models.UserModel) error {
	session := databases.GetMGO()
	defer session.Close()

	col := session.DB("test").C("users")
	u2 := uuid.NewV4()

	user.ID = u2.String()
	user.Token = "token"
	user.Stamp()
	err := col.Insert(user)
	if err != nil {
		log.Print("error insert", err)
	}
	return nil
}
