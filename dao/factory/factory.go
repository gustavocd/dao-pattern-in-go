package factory

import (
	"github.com/gustavocd/dao-pattern/dao/interfaces"
	"github.com/gustavocd/dao-pattern/dao/mysql"
	"log"
)

func FactoryDao(e string) interfaces.UserDao {
	var i interfaces.UserDao
	switch e {
	case "mysql":
		i = mysql.UserImplMysql{}
	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}
