package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func UpdatePessoas(c *gin.Context) {
	var pessoa models.Pessoa

	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(400, gin.H{"error": "Os dados  inceridos são inválidos"})
		return
	}

	pessoaRepo := &repository.UpdatePessoasRepository{}

	if err := pessoaRepo.UpdatePessoas(pessoa); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar pessoa, tente novamente"})
		return
	}
	c.JSON(200, gin.H{"message": "Pessoa atualizada com sucesso"})
}
