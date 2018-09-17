package model

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
)

type User struct {
	BaseFields
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Status   int    `json:"status,omitempty"`
	RoleId   int    `json:"role_id,omitempty"`
	TeamId   int    `json:"team_id,omitempty"`
}

func CheckUserPasswd(username string, passwd string) *User {
	obj := new(User)
	obj.Username = username
	obj.Password = passwd
	obj.Status = 1
	ExsitAndFirst(obj)
	return obj
}

func GetUserFromToken(c *gin.Context) *User {
	claims := jwt.ExtractClaims(c)
	user := new(User)
	ByID(user, int(claims["uid"].(float64)))
	return user
}

func (mod *User) Team() *Team {
	team := &Team{}
	ByID(team, mod.TeamId)
	return team
}