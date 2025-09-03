package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type DeletePessoasRepository struct {
}

func (r *DeletePessoasRepository) DeletePessoas(db *sql.DB, pessoa models.Pessoa) {
	fmt.Println("Tentando deletar com a Procedure")

	err := r.deletePessoasWhithProcedure(db, pessoa)
	if err == nil {
		fmt.Println("Procedure executada com sucesso")
		return
	}
	time.Sleep(10 * time.Second)

	fmt.Println("Procedure falhou, tentando com a Query")

	err = r.deletePessoasWhithQuery(db, pessoa)
	if err == nil {
		fmt.Println("Query executada com sucesso")
		return
	}

}
func (r *DeletePessoasRepository) deletePessoasWhithProcedure(db *sql.DB, pessoa models.Pessoa) error {

	_, err := db.Exec("CALL DeletePessoa(?, ?)",
		pessoa.Id,
		pessoa.Deletado_em)
	return err
}
func (r *DeletePessoasRepository) deletePessoasWhithQuery(db *sql.DB, pessoa models.Pessoa) error {

	query := `UPDATE pessoas SET deletado_em = ? WHERE id = ?`
	_, err := db.Exec(query,
		pessoa.Deletado_em,
		pessoa.Id)
	if err != nil {
		return err
	}
	return nil
}
