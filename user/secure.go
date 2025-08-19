package user

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("4P183ND4H4R453K0744")

func GenerateToken(input DataTokenInput) (string, error) {
	claim := jwt.MapClaims{}
	claim["Id_user"] = input.Id_user
	claim["Username"] = input.Username
	claim["Full_name"] = input.Full_name

	ttl := 100 * time.Hour
	claim["exp"] = time.Now().UTC().Add(ttl).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
