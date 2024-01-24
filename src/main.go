package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/src/app/config/awsgo"
	"gambituser/src/app/config/bd"
	"gambituser/src/app/constants"
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

	_, validParament := ValidParament()

	if !validParament {
		fmt.Println(constants.ErrorParament)
		err := errors.New(constants.ErrorParament)
		return events, err
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
		log.Fatalf(fmt.Sprintf("Error al leer el secret: %v", err.Error()))
		return events, err
	}

	return events, bd.SignUp(date)
}

func ValidParament() (string, bool) {
	return os.LookupEnv("SecretName")
}
