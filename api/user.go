package api

import (
	"github.com/beforesecond/gqlgen-todos/models"
)

const kindUser = "User"

// FindUser from datastore
func FindUser(username, password string) (*models.UserModel, error) {

	var user models.UserModel
	// if !user.ComparePassword(password) {
	// 	// wrong password return like user not found
	// 	return nil, nil
	// }
	user.ID = "1"
	return &user, nil
}

// SaveUser to datastore
func SaveUser(user *models.UserModel) error {
	// ctx, cancel := getContext()
	// defer cancel()
	// var err error
	// user.Stamp()
	// key := user.Key()
	// if key == nil {
	// 	key = datastore.IncompleteKey(kindUser, nil)
	// }
	// key, err = client.Put(ctx, key, user)
	// if err != nil {
	// 	return err
	// }
	// user.SetKey(key)
	return nil
}
