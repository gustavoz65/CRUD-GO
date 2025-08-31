package repository

import (
	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type DeletePessoasRepository struct {
}

func (r *DeletePessoasRepository) DeletePessoas(pessoa models.Pessoa) error {
	db := database.ConectarDB()
	defer db.Close()

	query := `UPDATE pessoas SET deletado_em = ? WHERE id = ?`
	_, err := db.Exec(query, pessoa.Deletado_em, pessoa.Id)
	if err != nil {
		return err
	}
	return nil
}
