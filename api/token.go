package api

import (
	"time"

	"github.com/beforesecond/gqlgen-todos/databases"
	"github.com/beforesecond/gqlgen-todos/models"
	"gopkg.in/mgo.v2/bson"
)

const kindToken = "Token"

// CreateToken save new token to database
func CreateToken(token string, userID string) error {
	session := databases.GetMGO()
	defer session.Close()

	col := session.DB("test").C("users")
	query := bson.M{
		"user.id": userID,
	}

	user := models.UserModel{}

	err := col.Find(query).One(&user)

	if err != nil {
		return err
	}
	user.Token = token
	user.Stamp()

	err = col.Update(query, user)

	if err != nil {
		return err
	}
	return nil
}
func getToken(token string) (*models.Token, error) {

	return nil, nil
}

func DeleteToken(token string) error {
	//tk, err := getToken(token)
	// if err != nil {
	// 	return err
	// }

	//defer cancel()
	return nil
}

// ValidateToken validate and update token last access timestamp
func ValidateToken(token string, userID string, expiresInFromLastAccess time.Duration) (bool, error) {
	// tk, err := getToken(token)
	// if err != nil {
	// 	return false, err
	// }
	// if tk == nil || tk.UserID != userID {
	// 	return false, nil
	// }
	// if time.Now().After(tk.LastAccessAt.Add(expiresInFromLastAccess)) {
	// 	// token expired
	// 	// remove expired token from database
	// 	go DeleteToken(token)
	// 	return false, nil
	// }
	// tk.Stamp()
	// go func(tk model.Token) {
	// 	ctx, cancel := getContext()
	// 	defer cancel()
	// 	client.Put(ctx, tk.Key(), &tk)
	// }(*tk)
	return true, nil
}
