package auth

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	jwt := NewJWt()
	token, err := jwt.buildToken(
		time.Now().Unix(),
		time.Now().Add(time.Second).Unix(),
		"jti",
		&BaseClaims{
			UserId:   "xxxxxx",
			UserName: "xxx",
		},
	)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(token)
	claims, err := jwt.parseToken(token)
	if err != nil {
		t.Log(err)
	}
	for k, v := range claims {
		switch k {
		case "aud", "exp", "jti", "iat", "iss", "nbf", "sub":
		// ignore the standard claims
		default:
			//b, _ := json.Marshal(v)
			//fmt.Println(string(b))
			var claims BaseClaims
			//json.Unmarshal(b, &claims)
			//fmt.Println(claims)
			copier.Copy(&v, &claims)
			fmt.Println(claims)
		}

	}
}
