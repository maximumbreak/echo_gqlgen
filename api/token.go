package api

import (
	"time"

	"github.com/beforesecond/gqlgen-todos/models"
)

const kindToken = "Token"

// CreateToken save new token to database
func CreateToken(token string, userID string) error {
	//  ctx, cancel := getContext()
	//  defer cancel()
	//  var err error
	//  tk := &model.Token{
	//   Token:  token,
	//   UserID: userID,
	//  }
	//  tk.Stamp()
	//  key := datastore.IncompleteKey(kindToken, nil)
	//  key, err = client.Put(ctx, key, tk)
	//  if err != nil {
	//   return err
	//  }
	//  tk.SetKey(key)
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
