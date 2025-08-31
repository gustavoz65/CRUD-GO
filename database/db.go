package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarDB() *sql.DB {
	db, err := sql.Open("mysql", "root:senha123@tcp(127.0.0.1:3306)/CRUDGio?parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conexão estabelecida com sucesso")
	}
	db.SetMaxOpenConns(25)                 // Máximo de conexões simultâneas
	db.SetMaxIdleConns(10)                 // Conexões que ficam "dormindo" esperando
	db.SetConnMaxLifetime(5 * time.Minute) // Tempo máximo de vida de uma conexão
	db.SetConnMaxIdleTime(3 * time.Minute) // Tempo máximo que uma conexão pode ficar ociosa
	return db
}
