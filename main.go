package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	gin.SetMode(gin.ReleaseMode)

	db := database.ConectarDB()

	defer db.Close()

}
