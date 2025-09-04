package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	// Rotas da API
	r.POST("/pessoas", controller.CreatePessoas)
	r.GET("/pessoas", controller.GetPessoas)
	r.GET("/pessoas/:id", controller.GetPessoaById)
	r.PUT("/pessoas/:id", controller.UpdatePessoas)
	r.DELETE("/pessoas/:id", controller.DeletePessoas)
	r.GET("/pessoas/:id/endereco", controller.GetCepPessoas)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
