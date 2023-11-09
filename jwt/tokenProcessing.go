package jwt

import (
	"errors"
	"strings"

	"github.com/brayanzuritadev/citas/db"
	"github.com/brayanzuritadev/citas/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func TokenProcess(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Incorrect token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {

		return myKey, nil
	})

	if err == nil {
		_, err = db.GetUser(claims.Email)
		if err == nil {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return &claims, true, IDUser, nil
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Invalid token")
	}

	return &claims, true, string(""), err
}
