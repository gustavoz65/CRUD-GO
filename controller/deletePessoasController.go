package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func DeletePessoas(c *gin.Context) {
	var pessoa models.Pessoa
	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(400, gin.H{"error": "Dados Invalidos"})
		return
	}

	pessoaRepo := &repository.DeletePessoasRepository{}

	if err := pessoaRepo.DeletePessoas(pessoa); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao deletar pessoa, tente novamente"})
		return
	}
	c.JSON(200, gin.H{"message": "Pessoa deletada com sucesso", "id": pessoa.Id})
}
