package middleware

import (
	"github.com/appleboy/gin-jwt"
	"time"
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"apibuilder-server/handler/endpoint"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func InitJWT() {
	// the jwt middleware
	AuthMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "API Builder tool",
		Key:        []byte("build restful api"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (interface{}, bool) {
			user := model.CheckUserPasswd(userId, password)
			if user.ID > 0 {
				return user, true
			}
			return nil, false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			claims  := make(map[string]interface{})
			if data != nil{
				user := data.(*model.User)
				claims["uid"] = user.ID
				claims["email"] = user.Email
				claims["phone"] = user.Phone
			}
			return claims
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			return true

			//if v, ok := user.(string); ok && v == "admin" {
			//	return true
			//}
			//
			//return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			panic(endpoint.NewControllerError(code, "unauthorized", "Unauthorized", message))
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}

func IsAdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//claims := jwt.ExtractClaims(c)
		//user := model.GetUserFromToken(c)
		c.Next()
	}
}
