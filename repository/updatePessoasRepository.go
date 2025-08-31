package repository

import (
	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type UpdatePessoasRepository struct {
}

func (r *UpdatePessoasRepository) UpdatePessoas(pessoa models.Pessoa) error {

	db := database.ConectarDB()
	defer db.Close()

	query := `UPDATE pessoas SET 
          nome = ?, descricao = ?, ativo = ?, altura_metros = ?, 
          nascimento = ?, cep = ?, atualizado_em = ? 
          WHERE id = ?`
	_, err := db.Exec(query, pessoa.Nome, pessoa.Descricao, pessoa.Ativo, pessoa.Altura_metros, pessoa.Nascimento, pessoa.Cep, pessoa.Atualizado_em, pessoa.Id)
	if err != nil {
		return err
	}
	return nil

}
