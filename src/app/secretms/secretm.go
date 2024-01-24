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

	log.Printf("---Solicitud de Nombre Secreto: %s---", nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		return dateSecret, fmt.Errorf("Error al obtener el secreto: %w", err)
	}

	log.Printf("Valor del secreto: %v", *key.SecretString)

	if key.SecretString == nil {
		return dateSecret, fmt.Errorf("El secreto no contiene una cadena válida")
	}

	log.Printf("Clave: %v", key)

	err = json.Unmarshal([]byte(*key.SecretString), &dateSecret)
	if err != nil {
		return dateSecret, fmt.Errorf("%v= %v", constants.DeserializationError, err)
	}

	return dateSecret, nil
}
