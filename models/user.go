package models

// type UserModel struct {
// 	Text   string `json:"text"`
// 	UserID string `json:"userId"`
// }

type UserModel struct {
	User
	HasPassword
	HasTimestamp
}
