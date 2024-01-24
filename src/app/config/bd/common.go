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
	if err != nil {
		return fmt.Errorf("Error al leer el secreto: %w", err)
	}
	return err
}

func DbConnect() (*sql.DB, error) {
	connStr := ConnStr(secret)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error al abrir la conexión a la base de datos: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error al recibir respuesta: %w", err)
	}

	log.Print("Conexión exitosa de DB.")
	return db, nil
}

func ConnStr(keys models.SecretRds) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		keys.UserName, keys.Password, keys.Host, "gambit")
}
