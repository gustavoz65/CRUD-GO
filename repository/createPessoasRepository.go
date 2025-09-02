package repository

import (
	"database/sql"
	"fmt"

	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type CreatePessoasRepository struct {
}

func (r *CreatePessoasRepository) CreatePessoas(pessoa models.Pessoa) error {
	db := database.ConectarDB()
	defer db.Close()

	fmt.Println("Tentando criar pessoa com stored procedure...")
	err := r.createPessoaWithProcedure(db, pessoa)
	if err == nil {
		fmt.Println("Sucesso com stored procedure!")
		return nil
	}

	fmt.Println("Falha na stored procedure:", err)
	fmt.Println("Tentando com query direta...")

	err = r.createPessoaWithQuery(db, pessoa)
	if err != nil {
		fmt.Println("Falha também na query direta:", err)
		return err
	}

	fmt.Println("Sucesso com query direta!")
	return nil
}

func (r *CreatePessoasRepository) createPessoaWithProcedure(db *sql.DB, pessoa models.Pessoa) error {
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

func (r *CreatePessoasRepository) createPessoaWithQuery(db *sql.DB, pessoa models.Pessoa) error {
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
