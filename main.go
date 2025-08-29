package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := conectarDB()
	defer db.Close()

	fmt.Println("Conexão estabelecida com sucesso")

	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)
}
