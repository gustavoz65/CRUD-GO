package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/controller"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/createPessoa", controller.CreatePessoas)
	r.PUT("/updatePessoa", controller.UpdatePessoas)
	r.Run(":8080")
}
