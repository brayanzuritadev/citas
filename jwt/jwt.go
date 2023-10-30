package jwt

import (
	"context"
	"time"

	"github.com/brayanzuritadev/citas/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, user models.User) (string, error) {

	jwtSign := ctx.Value(models.Key("jwtsign")).(string)
	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":     user.Email,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"datebirth": user.DateBirth,
		"profile":   user.Profile,
		"address":   user.Address,
		"web":       user.Web,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
