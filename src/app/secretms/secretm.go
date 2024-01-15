package secretms

import (
	"context"
	"encoding/json"
	"fmt"
	"gambituser/src/app/config/awsgo"
	"gambituser/src/app/constants"
	"gambituser/src/app/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
)

func GetSecret(ctx context.Context, nameSecret string) (models.SecretRds, error) {
	var dateSecret models.SecretRds
	log.Println("--- Solicitud de Nombre Secreto ---")

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		msnError := fmt.Errorf("%v", err.Error())
		return dateSecret, msnError
	}

	err = json.Unmarshal([]byte(*key.SecretString), &dateSecret)
	if err != nil {
		msnError := fmt.Errorf("%v= %v", constants.DeserializationError, err)
		return dateSecret, msnError
	}

	return dateSecret, nil
}
