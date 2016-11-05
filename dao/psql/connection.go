package psql

import (
	"database/sql"
	"github.com/gustavocd/dao-pattern/utilities"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func get() *sql.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return db
}
