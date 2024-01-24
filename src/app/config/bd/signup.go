package bd

import (
	"fmt"
	"gambituser/src/app/models"
	"gambituser/src/app/tools"
	"log"
)

func SignUp(sign models.SignUp) error {
	log.Print("Registro iniciado")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sign.UserEmail + "','" + sign.UserUUID + "','" + tools.DateMySql() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		return err
	}
	return nil
}
