package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func GetPessoas(c *gin.Context) {
	var pessoa models.Pessoa

	if err := c.ShouldBindJSON(&pessoa); err != nil {
		c.JSON(400, gin.H{"error": "dados Invalidos"})
		return
	}

	pessoaRepo := &repository.GetPessoasRepository{}
	if err := pessoaRepo.GetPessoas(pessoa); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao Buscar Pessoas no Banco de Dados"})
		return
	}
	c.JSON(200, gin.H{"message": "Pessoa achada com sucesso", "nome": pessoa.Nome})
}
