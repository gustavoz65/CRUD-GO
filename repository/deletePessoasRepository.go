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

	query := `DELETE FROM pessoas WHERE id = ?`
	_, err := db.Exec(query, pessoa.Id)
	if err != nil {
		return err
	}
	return nil
}
