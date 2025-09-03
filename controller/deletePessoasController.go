package controller

import (
	"database/sql"
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

	// Get DB from Gin context (assuming it's set as "db")
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "Erro ao obter conexão com banco de dados"})
		return
	}

	pessoaRepo.DeletePessoas(db, models.Pessoa{Id: id})
	c.JSON(200, gin.H{"message": "Pessoa deletada com sucesso", "id": id})
}
