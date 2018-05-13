package main

import (
	"fmt"
	"log"

	"github.com/gustavocd/dao-pattern-in-go/dao/factory"
	"github.com/gustavocd/dao-pattern-in-go/models"
	"github.com/gustavocd/dao-pattern-in-go/utilities"
)

func main() {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	userDao := factory.FactoryDao(config.Engine)

	user := models.User{}
	fmt.Println("Digite su nombre:")
	fmt.Scan(&user.FirstName)
	fmt.Println("Digite su apellido:")
	fmt.Scan(&user.LastName)
	fmt.Println("Digite su correo:")
	fmt.Scan(&user.Email)

	err = userDao.Create(&user)

	if err != nil {
		log.Fatal(err)
		return
	}

}
