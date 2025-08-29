package repository

import (
	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type CreatePessoasRepository struct {
}

func (r *CreatePessoasRepository) CreatePessoas(pessoa models.Pessoa) error {

	db := database.ConectarDB()
	defer db.Close()

	query := `INSERT INTO pessoas (nome, descricao) VALUES (?, ?)`
	_, err := db.Exec(query, pessoa.Nome, pessoa.Descricao)
	if err != nil {
		return err
	}
	return nil
}
