package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

const (
	bearer        = "Bearer"
	bearerFormat  = "Bearer %s"
	authorization = "Authorization"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

type JWT struct {
	SignKey []byte
}

type BaseClaims struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func NewJWt() *JWT {
	return &JWT{
		SignKey: []byte("xxxxxxx"),
	}
}

func (j *JWT) buildToken(iat, exp int64, jti string, claims *BaseClaims) (string, error) {
	c := make(jwt.MapClaims)
	c["exp"] = iat + exp
	c["iat"] = iat
	c["jti"] = jti
	c["claims"] = claims
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = c
	return t.SignedString(j.SignKey)
}

func (j *JWT) parseToken(token string) (jwt.MapClaims, error) {
	var (
		t   *jwt.Token
		err error
	)

	t, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if t != nil {
		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
