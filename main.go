package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/brayanzuritadev/citas/awsgo"
	"github.com/brayanzuritadev/citas/db"
	"github.com/brayanzuritadev/citas/handlers"
	"github.com/brayanzuritadev/citas/models"
	"github.com/brayanzuritadev/citas/secretmanager"
)

func main() {
	awsgo.InitAWS()
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	if !ValidParameters() {
		return ErrorResponse(400, "Error in environment variables. Must include 'SecretName', 'BucketName', 'UrlPrefix'")
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		return ErrorResponse(400, "Error reading Secret "+err.Error())
	}

	path := strings.Replace(request.PathParameters["dailygo"], os.Getenv("UrlPrefix"), "", -1)
	fmt.Println(path)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("username"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//review connection db
	err = db.ConnetionDb(awsgo.Ctx)
	defer db.CloseConnection()

	if err != nil {
		return ErrorResponse(500, "Error connecting to database "+err.Error())
	}

	responseApi := handlers.Handlers(awsgo.Ctx, request)
	if responseApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: responseApi.Status,
			Body:       responseApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return responseApi.CustomResp, nil
	}
}

func ErrorResponse(statusCode int, message string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       message,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func ValidParameters() bool {
	_, getParameter := os.LookupEnv("SecretName")
	if !getParameter {
		return getParameter
	}

	_, getParameter = os.LookupEnv("BucketName")
	if !getParameter {
		return getParameter
	}

	_, getParameter = os.LookupEnv("UrlPrefix")
	if !getParameter {
		return getParameter
	}
	return getParameter
}
