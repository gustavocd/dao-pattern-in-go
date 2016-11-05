package main

import (
	"github.com/gustavocd/dao-pattern/utilities"
	"log"
	"github.com/gustavocd/dao-pattern/dao/factory"
	"github.com/gustavocd/dao-pattern/models"
	"fmt"
)

func main()  {
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

	if err != nil{
		log.Fatal(err)
		return
	}


}
