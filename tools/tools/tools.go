package tools

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
)

// JWTCreateAuthorizationBy32
// @Desc：生成32字节的密钥
// @param：mapClaims
// @return：string
// @return：error
func JWTCreateAuthorizationBy32(mapClaims jwt.MapClaims) (string, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	jwtStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return jwtStr, nil
}
