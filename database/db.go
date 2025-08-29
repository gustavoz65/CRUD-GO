package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func conectarDB() *sql.DB {
	db, err := sql.Open("mysql", "root:senha123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexão estabelecida com sucesso")
	}
	return db
}
