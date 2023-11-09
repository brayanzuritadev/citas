package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/brayanzuritadev/citas/db"
	"github.com/brayanzuritadev/citas/jwt"
	"github.com/brayanzuritadev/citas/models"
)

func Login(ctx context.Context) models.ResponseApi {
	var user models.User
	var responseApi models.ResponseApi

	responseApi.Status = 400

	if value, ok := ctx.Value(models.Key("body")).(string); ok {
		fmt.Println(value)
	} else {
		fmt.Println("son tonteras")
	}

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		fmt.Println("s")
		responseApi.Message = "Email or Password invalid" + err.Error()
		return responseApi
	}

	if len(user.Email) == 0 {
		responseApi.Message = "The Email is required"
		return responseApi
	}

	userData, exist := db.LoginTry(user.Email, user.Password)
	if !exist {
		responseApi.Message = "Email and Password invalid "
		return responseApi
	}

	jwtKey, err := jwt.GenerateJWT(ctx, userData)

	if err != nil {
		responseApi.Message = "An error occurred while trying to generate the token"
		return responseApi
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		responseApi.Message = "An error occurred while trying to generate the JSON > " + err2.Error()
		return responseApi
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	responseApi.Status = 200
	responseApi.Message = string(token)
	responseApi.CustomResp = res

	return responseApi
}
