package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Details struct {
	SecretKey string
	UserId    int
	UniqueId  string
	Duration  time.Duration
}

func Generate(td Details) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = td.UserId
	claims["unique_id"] = td.UniqueId
	claims["exp"] = time.Now().Add(td.Duration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(td.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Parse(secretKey, s string) (*Details, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(s, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unknown Method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Unauthorized")
	}

	td := Details{
		UserId:   int(claims["user_id"].(float64)),
		UniqueId: claims["unique_id"].(string),
	}

	return &td, nil
}
