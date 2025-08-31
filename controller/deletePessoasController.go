package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func DeletePessoas(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID invalido"})
		return
	}

	pessoaRepo := &repository.DeletePessoasRepository{}

	if err := pessoaRepo.DeletePessoas(models.Pessoa{Id: id}); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao deletar pessoa, tente novamente"})
		return
	}
	c.JSON(200, gin.H{"message": "Pessoa deletada com sucesso", "id": id})
}
