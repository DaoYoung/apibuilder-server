package model

import (
	"time"
	"hlj-comet/app"
)

type User struct {
	Id        int       `json:"id"`
	Kind      int       `json:"kind"`
	Nick      string    `json:"nick"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Cid       int       `json:"cid"`
	UserToken string    `json:"-"`
	CreatedAt time.Time `json:"created_at" time_format:"sql_datetime"`
}

func GetUserById(id int) *User {
	user := User{}
	if app.WedDb.Where("id = ?", id).Last(&user).Error == nil {
		return &user
	} else {
		return nil
	}
}

func GetUserByToken(token string) *User {
	user := User{}
	if app.WedDb.Where("user_token = ?", token).Last(&user).Error == nil {
		return &user
	} else {
		return nil
	}
}

func (User) TableName() string {
	return "users"
}