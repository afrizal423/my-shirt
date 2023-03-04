package middleware

import (
	"kemejaku/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	UserId int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string //ambil dari config json
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.ErrorResponse(c, http.StatusForbidden, e.Error(), e)
		}),
	}
}

func (configJwt ConfigJWT) GenerateToken(userId int) string {
	claims := JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJwt.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(configJwt.SecretJWT))

	return token
}

func GetUserId(c echo.Context) int {
	if temp := c.Get("user"); temp != nil {
		user := temp.(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		return int(claims.UserId)
	}
	return 0
}
