package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-NBS-GO/repository"
)

// GetCepPessoas godoc
// @Summary     Busca endereço por ID da pessoa
// @Description Busca o CEP de uma pessoa no banco de dados e consulta o ViaCEP para retornar informações completas do endereço
// @Tags        Pessoas
// @Accept      json
// @Produce     json
// @Param       id path int true "ID da pessoa"
// @Success     200 {object} map[string]interface{} "Endereço encontrado com sucesso"
// @Failure     400 {object} map[string]string "ID inválido ou pessoa sem CEP cadastrado"
// @Failure     404 {object} map[string]string "Pessoa não encontrada"
// @Failure     500 {object} map[string]string "Erro interno do servidor"
// @Router      /pessoas/{id}/endereco [get]
func GetCepPessoas(c *gin.Context) {
	// 1. Extrai o ID da pessoa da URL (parâmetro :id)
	idParam := c.Param("id")

	// 2. Converte o ID de string para int
	pessoaId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// 3. Cria uma instância do repository
	pessoaRepo := &repository.GetCepPessoasRepository{}

	// 4. Busca o CEP da pessoa pelo ID e consulta o ViaCEP
	endereco, err := pessoaRepo.GetCepPessoasById(pessoaId)
	if err != nil {
		// 5. Retorna erro específico baseado no tipo de erro
		if err.Error() == "pessoa não encontrada" {
			c.JSON(404, gin.H{"error": "Pessoa não encontrada"})
			return
		}
		if err.Error() == "pessoa não possui CEP cadastrado" {
			c.JSON(400, gin.H{"error": "Pessoa não possui CEP cadastrado"})
			return
		}
		c.JSON(500, gin.H{"error": "Erro ao buscar o endereço: " + err.Error()})
		return
	}

	// 6. Retorna o endereço encontrado
	c.JSON(http.StatusOK, gin.H{
		"pessoa_id": pessoaId,
		"endereco":  endereco,
	})
}
