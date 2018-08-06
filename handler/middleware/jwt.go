package middleware

import (
	"github.com/appleboy/gin-jwt"
	"time"
	"github.com/gin-gonic/gin"
	"apibuilder-server/model"
	"apibuilder-server/handler/endpoint"
	"net/http"
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
func  RefreshHandler(c *gin.Context) {
	token, _ := AuthMiddleware.parseToken(c)
	claims := token.Claims.(jwt.MapClaims)

	origIat := int64(claims["orig_iat"].(float64))

	if origIat < AuthMiddleware.TimeFunc().Add(-AuthMiddleware.MaxRefresh).Unix() {
		AuthMiddleware.unauthorized(c, http.StatusUnauthorized, AuthMiddleware.HTTPStatusMessageFunc(ErrExpiredToken, c))
		return
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(AuthMiddleware.SigningAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire := AuthMiddleware.TimeFunc().Add(AuthMiddleware.Timeout)
	newClaims["id"] = claims["id"]
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = AuthMiddleware.TimeFunc().Unix()
	tokenString, err := AuthMiddleware.signedString(newToken)

	if err != nil {
		AuthMiddleware.unauthorized(c, http.StatusUnauthorized, AuthMiddleware.HTTPStatusMessageFunc(ErrFailedTokenCreation, c))
		return
	}

	// set cookie
	if AuthMiddleware.SendCookie {
		maxage := int(expire.Unix() - time.Now().Unix())
		c.SetCookie(
			"JWTToken",
			tokenString,
			maxage,
			"/",
			"",
			AuthMiddleware.SecureCookie,
			true,
		)
	}

	AuthMiddleware.RefreshResponse(c, http.StatusOK, tokenString, expire)
}
