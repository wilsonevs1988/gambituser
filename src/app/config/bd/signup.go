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

	insert := "INSERT INTO users (User_Email,User_FirstName,User_LastName,User_Status,User_DateAdd) VALUES ('" + sign.UserEmail + "', '" + sign.UserUUID + "', '" + tools.DateMySql() + "')"
	fmt.Println(insert)

	_, err = Db.Exec(insert)
	if err != nil {
		return err
	}
	return nil
}
