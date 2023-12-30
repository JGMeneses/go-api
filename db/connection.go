package db

import (
	"database/sql"
	"fmt"

	"github.com/JGMeneses/go-crud/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host%s port%s user%s password%s dbname%s sslmode=disable",
		conf.Host, conf.Port, conf.Pass, conf.Database)

	conn, err := sql.Open("postgress", sc)

	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}
