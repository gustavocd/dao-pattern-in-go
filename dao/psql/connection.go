package psql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gustavocd/dao-pattern-in-go/utilities"
	_ "github.com/lib/pq"
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
