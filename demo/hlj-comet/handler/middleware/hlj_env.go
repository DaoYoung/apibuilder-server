package middleware

import (
	"hlj-rest/rest"
	"hlj-comet/model"
)

func RemoteUser(r *rest.Request) *model.User {
	remoteUser, ok := r.Env["REMOTE_USER"]
	if !ok {
		return nil
	}
	user, ok := remoteUser.(*model.User)
	if !ok || user == nil {
		return nil
	}

	return user
}

func RemoteRoleUser(r *rest.Request) *model.User {
	remoteUser, ok := r.Env["REMOTE_ROLE_USER"]
	if !ok {
		return nil
	}
	user, ok := remoteUser.(*model.User)
	if !ok || user == nil {
		return nil
	}

	return user
}