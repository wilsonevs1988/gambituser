package bd

import (
	"database/sql"
	"fmt"
	"gambituser/src/app/config/awsgo"
	"gambituser/src/app/models"
	"gambituser/src/app/secretms"
	"log"
	"os"
)

var secret models.SecretRds
var err error
var Db *sql.DB

func ReadSecret() error {
	secret, err = secretms.GetSecret(awsgo.Ctx, os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(secret))
	if err != nil {
		log.Fatalln("Error al abrir la conexión a la base de datos: ", err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalln("Error al recibir respuesta: ", err.Error())
		return err
	}
	log.Print("Conexión exitosa de DB.")
	return err
}

func ConnStr(keys models.SecretRds) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = keys.UserName
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gambit"

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser, authToken, dbEndpoint, dbName)
}
