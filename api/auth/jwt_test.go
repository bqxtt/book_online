package auth

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := GenerateToken(10175101201, "admin")
	if err == nil {
		fmt.Println(token)
	}
	claims, err := ParseToken(token)
	if err == nil {
		fmt.Println(claims)
	}
}
