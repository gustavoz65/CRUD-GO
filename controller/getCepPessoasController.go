package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func GetCepPessoas(c *gin.Context) {
	cep := c.Param("cep")
	pessoaRepo := &repository.GetCepPessoasRepository{}

	endereco, err := pessoaRepo.GetCepPessoas(cep)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar o cep, tente novamente"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"endereço": endereco})
}
