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
	lambda.Start(ExecuteLambda)
	defer db.CloseConnection()
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InitAWS()

	if !ValidParameters() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in environment variables. Must include 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error to read the secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["dailygo"], os.Getenv("UrlPrefix"), "", -1)
	fmt.Println(path + "este es el path")
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

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connect to the db " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
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
