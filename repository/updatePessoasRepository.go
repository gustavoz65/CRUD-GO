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

	query := `UPDATE pessoas SET nome = ?, descricao = ? WHERE id = ?`
	_, err := db.Exec(query, pessoa.Nome, pessoa.Descricao, pessoa.Id)
	if err != nil {
		return err
	}
	return nil

}
