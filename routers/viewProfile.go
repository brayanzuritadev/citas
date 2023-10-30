package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/brayanzuritadev/citas/db"
	"github.com/brayanzuritadev/citas/models"
)

func ViewProfile(request events.APIGatewayProxyRequest) models.ResponseApi {

	var r models.ResponseApi
	r.Status = 400

	fmt.Println("View Profile")

	Id := request.QueryStringParameters["id"]
	if len(Id) < 1 {
		r.Message = "Id is required"
		return r
	}

	profile, err := db.SearchProfile(Id)

	if err != nil {
		r.Message = "An error occurred while trying to search the user "
		return r
	}

	respJson, err := json.Marshal(profile)
	if err != nil {
		r.Status = 500
		r.Message = "An error occurred while trying to generate the JSON > " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
