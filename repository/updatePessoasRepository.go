package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type UpdatePessoasRepository struct {
}

func (r *UpdatePessoasRepository) UpdatePessoas(pessoa models.Pessoa) error {
	db := database.ConectarDB()
	defer db.Close()

	time.Sleep(10 * time.Second)

	fmt.Println("Tentando com a Procedure")

	err := r.updatePessoaWhithProcedure(db, pessoa)
	if err == nil {
		fmt.Println("Procedure executada com sucesso")
		return nil
	}
	fmt.Println("Procedure falhou, tentando com a Query")

	time.Sleep(10 * time.Second)

	err = r.updatePessoaWhithQuery(db, pessoa)
	if err == nil {
		fmt.Println("Query executada com sucesso")
		return nil
	}
	return err

}

func (r *UpdatePessoasRepository) updatePessoaWhithProcedure(db *sql.DB, pessoa models.Pessoa) error {
	query := `CALL sp_atualizar_pessoa (?,?,?,?,?,?,?,?)`

	_, err := db.Exec(
		query,
		pessoa.Id,
		pessoa.Nome,
		pessoa.Descricao,
		pessoa.Ativo,
		pessoa.Altura_metros,
		pessoa.Nascimento,
		pessoa.Cep,
		pessoa.Atualizado_em,
	)
	return err
}

func (r *UpdatePessoasRepository) updatePessoaWhithQuery(db *sql.DB, pessoa models.Pessoa) error {
	query := `UPDATE pessoas SET 
	nome = ?,
	descricao = ?,
	ativo = ?,
	altura_metros = ?,
	`
	_, err := db.Exec(
		query,
		pessoa.Nome,
		pessoa.Descricao,
		pessoa.Ativo,
		pessoa.Altura_metros,
		pessoa.Nascimento,
		pessoa.Cep,
		pessoa.Atualizado_em,
	)
	return err
}
