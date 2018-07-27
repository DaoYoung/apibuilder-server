// 婚礼纪用户授权中间件
//
// liamylian
// 2018/03/09

package middleware

import (
	"hlj-rest/rest"
	"log"
	"net/http"
	"net/url"
)

type HljAuthMiddleware struct {
	Authenticator func(userToken string) (user interface{}, ok bool)

	Unauthenticated func(writer rest.ResponseWriter, request *rest.Request)
}

var unauthenticated = func(writer rest.ResponseWriter, request *rest.Request) {
	rest.Error(writer, http.StatusUnauthorized, "Not Authorized", http.StatusUnauthorized)
}

func (mw *HljAuthMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	if mw.Authenticator == nil {
		log.Fatal("Authenticator is required")
	}

	return func(writer rest.ResponseWriter, request *rest.Request) {
		loginAs := request.Header.Get("login-as")
		if loginAs == "" {
			loginAs = request.QueryParam("login_as").String()
		}

		roleTokenKey := "role-token"
		if loginAs == "main" {
			roleTokenKey = "main-role-token"
		}
		roleUserToken := request.URL.Query().Get(roleTokenKey)
		if roleUserToken == "" {
			roleUserToken = request.Header.Get(roleTokenKey)
		}
		if roleUserToken == "" {
			if cookie, err := request.Cookie(roleTokenKey); err == nil {
				roleUserToken = cookie.Value
			}
		}

		userToken := request.URL.Query().Get("token")
		if userToken == "" {
			userToken = request.Header.Get("token")
		}
		if userToken == "" {
			if cookie, err := request.Cookie("token"); err == nil {
				userToken = cookie.Value
			}
		}
		if userToken == "" {
			tokenKey := "user"
			if loginAs == "main" {
				tokenKey = "main_user"
			}
			if cookie, err := request.Cookie(tokenKey); err == nil {
				userToken = cookie.Value
			}
		}

		if token, err := url.QueryUnescape(userToken); err == nil {
			userToken = token
		}
		if token, err := url.QueryUnescape(roleUserToken); err == nil {
			roleUserToken = token
		}

		user, userOk := mw.Authenticator(userToken)
		roleUser, roleUserOk := mw.Authenticator(roleUserToken)
		if !userOk && !roleUserOk {
			if mw.Unauthenticated != nil {
				mw.Unauthenticated(writer, request)
			} else {
				unauthenticated(writer, request)
			}
			return
		}
		if roleUser == nil {
			roleUser = user
		} else if user == nil {
			user = roleUser
		}

		request.Env["REMOTE_USER"] = user
		request.Env["REMOTE_ROLE_USER"] = roleUser
		handler(writer, request)
	}
}