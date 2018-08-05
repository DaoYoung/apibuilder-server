package model

type User struct {
	BaseFields
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
	RoleId   int    `json:"role_id"`
}

func CheckUserPasswd(username string, passwd string) *User {
	obj := new(User)
	obj.Username = username
	obj.Password = passwd
	obj.Status = 1
	ExsitAndFirst(obj)
	return obj
}