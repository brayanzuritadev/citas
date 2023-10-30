package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/brayanzuritadev/citas/jwt"
	"github.com/brayanzuritadev/citas/models"
	"github.com/brayanzuritadev/citas/routers"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResponseApi {
	fmt.Println("Processssss " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.ResponseApi

	r.Status = 400

	isOk, statusCode, msg, _ := validationAuthorization(ctx, request)

	if !isOk {
		fmt.Println("Process no bien ")
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		case "login":
			return routers.Login(ctx)
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "viewprofile":
			return routers.ViewProfile(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Method invalid"
	return r
}

func validationAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token required", models.Claim{}
	}

	claim, allOk, msg, err := jwt.TokenProcess(token, ctx.Value(models.Key("jwtsign")).(string))

	if !allOk {
		if err != nil {
			fmt.Println("Error token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error token " + msg)
			return false, 401, err.Error(), models.Claim{}
		}
	}

	return true, 200, msg, *claim
}
