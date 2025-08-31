package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/controller"
)

func SetupRoutes(r *gin.Engine) {
	// Rotas para pessoas (padrão REST)
	r.POST("/pessoas", controller.CreatePessoas)
	r.GET("/pessoas", controller.GetPessoas)
	r.GET("/pessoas/:id", controller.GetPessoaById)
	r.PUT("/pessoas/:id", controller.UpdatePessoas)
	r.DELETE("/pessoas/:id", controller.DeletePessoas)
}
