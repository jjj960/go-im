package ctxdata

import "github.com/golang-jwt/jwt"

const Identify = "hysgo-im"

func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds //当前时间+过期时间
	claims["iat"] = iat           //生成时间
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
