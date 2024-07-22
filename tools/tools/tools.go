package tools

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net"
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

// GetServerIP
// @Desc：获取服务IP
// @return：*net.UDPAddr
// @return：error
func GetServerIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	if err != nil {
		return "", err
	}
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	ip := fmt.Sprintf("%s", ipAddress.IP.String())
	return ip, nil
}
