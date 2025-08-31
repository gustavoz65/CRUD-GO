package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

func GetPessoas(c *gin.Context) {
	pessoaRepo := &repository.GetPessoasRepository{}

	pessoas, err := pessoaRepo.GetPessoas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pessoas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pessoas": pessoas})
}

func GetPessoaById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pessoaRepo := &repository.GetPessoasRepository{}
	pessoa, err := pessoaRepo.GetPessoaById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pessoa": pessoa})
}
