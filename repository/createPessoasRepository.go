package repository

import (
	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type CreatePessoasRepository struct {
}

func (r *CreatePessoasRepository) CreatePessoas(pessoa models.Pessoa, useStoredProcedure bool) error {
	db := database.ConectarDB()
	defer db.Close()

	if useStoredProcedure {
		sp := `CALL sp_criar_pessoa(?, ?, ?, ?, ?, ?)`
		_, err := db.Exec(sp,
			pessoa.Nome,
			pessoa.Descricao,
			pessoa.Ativo,
			pessoa.Altura_metros,
			pessoa.Nascimento,
			pessoa.Cep)
		return err
	}

	query := `INSERT INTO pessoas (nome, descricao, ativo, altura_metros, nascimento, cep) 
              VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query,
		pessoa.Nome,
		pessoa.Descricao,
		pessoa.Ativo,
		pessoa.Altura_metros,
		pessoa.Nascimento,
		pessoa.Cep)

	return err
}
