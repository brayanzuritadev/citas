package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/brayanzuritadev/citas/db"
	"github.com/brayanzuritadev/citas/models"
)

func Register(ctx context.Context) models.ResponseApi {
	var user models.User
	var response models.ResponseApi

	response.Status = 400

	fmt.Println("Register endpoint")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		response.Message = err.Error()
		fmt.Println(response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "You must specify the email address"
		fmt.Println(response.Message)
		return response
	}

	if len(user.Password) < 6 {
		response.Message = "You must specify a password longer than 5 characters"
		fmt.Println(response.Message)
		return response
	}

	_, find := db.GetUser(user.Email)

	if find {
		response.Message = "There is already a user with this email"
		fmt.Println(response.Message)
		return response
	}

	_, status, err := db.InsertUser(user)
	if err != nil {
		response.Message = "Error occurred while trying to register the user"
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Message = "Failed to register user"
		fmt.Println(response.Message)
		return response
	}

	response.Status = 200
	response.Message = "Register OK"
	fmt.Println(response.Message)
	return response
}
