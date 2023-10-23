package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/brayanzuritadev/citas/awsgo"
	"github.com/brayanzuritadev/citas/models"
)

func GetSecret(secretName string) (models.Secret, error){
	var dateSecret models.Secret
	fmt.Println("> get secret " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dateSecret, err
	}
	
	json.Unmarshal([]byte(*key.SecretString), &dateSecret)
	fmt.Println("> Red Secret Ok " + secretName)
	return dateSecret, nil
}