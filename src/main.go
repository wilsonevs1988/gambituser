package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/src/app/config/awsgo"
	"gambituser/src/app/config/bd"
	"gambituser/src/app/models"
	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func main() {
	lambda.Start(MyLambda)
}

func MyLambda(ctx context.Context, events events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAws()

	validParament, secretName := ValidParament()

	if !validParament {
		log.Printf("Error: Parámetro '%s' no válido", secretName)
		return events, errors.New("Error de parámetro")
	}

	var date models.SignUp
	for row, att := range events.Request.UserAttributes {
		switch row {
		case "email":
			date.UserEmail = att
			log.Print(fmt.Sprintf("Email: %v", date.UserEmail))
		case "sub":
			date.UserUUID = att
			log.Print(fmt.Sprintf("Sub: %v", date.UserUUID))
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		log.Printf("Error al leer el secreto: %v", err)
		return events, err
	}

	return events, bd.SignUp(date)
}

func ValidParament() (bool, string) {
	secretName, validParament := os.LookupEnv("SecretName")
	return validParament, secretName
}
