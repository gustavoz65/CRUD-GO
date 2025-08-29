package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func CreatePessoas(c *gin.Context) {
	var pessoa models.Pessoa
	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(400, gin.H{"error": "dados Invalidos"})
		return
	}

	pessoaRepo := &repository.CreatePessoasRepository{}

	if err := pessoaRepo.CreatePessoas(pessoa); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao criar pessoas, tente novamente"})
		return
	}

	c.JSON(200, gin.H{"message": "Pessoa criada com sucesso"})
}
