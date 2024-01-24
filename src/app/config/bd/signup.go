package bd

import (
	"fmt"
	"gambituser/src/app/models"
	"gambituser/src/app/tools"
	"log"
)

func SignUp(sign models.SignUp) error {
	log.Print("Registro iniciado")

	db, err := DbConnect()
	if err != nil {
		return fmt.Errorf("Error al conectar a la base de datos: %w", err)
	}
	defer db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)"
	fmt.Println(sentencia)

	_, err = db.Exec(sentencia, sign.UserEmail, sign.UserUUID, tools.DateMySql())
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %w", err)
	}

	return nil
}
