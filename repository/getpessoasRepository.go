package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gustavoz65/CRUD-NBS-GO/database"
	"github.com/gustavoz65/CRUD-NBS-GO/models"
)

type GetPessoasRepository struct {
}

// GetPessoas retorna todas as pessoas ativas
func (r *GetPessoasRepository) GetPessoas() ([]models.Pessoa, error) {
	fmt.Println("=== INICIANDO BUSCA DE PESSOAS ===")

	db := database.ConectarDB()
	defer db.Close()

	fmt.Println("Conexão OK")

	query := `SELECT id, nome, descricao, ativo, altura_metros, nascimento, cep, 
			  deletado_em, atualizado_em FROM pessoas WHERE deletado_em IS NULL`

	fmt.Println("Query:", query)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERRO NA QUERY:", err)
		return nil, err
	}
	defer rows.Close()

	fmt.Println("Query executada com sucesso")

	var pessoas []models.Pessoa
	for rows.Next() {
		fmt.Println("Processando linha...")
		var pessoa models.Pessoa
		var deletadoEm, atualizadoEm sql.NullTime

		err := rows.Scan(
			&pessoa.Id,
			&pessoa.Nome,
			&pessoa.Descricao,
			&pessoa.Ativo,
			&pessoa.Altura_metros,
			&pessoa.Nascimento,
			&pessoa.Cep,
			&deletadoEm,
			&atualizadoEm,
		)
		if err != nil {
			fmt.Println("ERRO NO SCAN:", err)
			return nil, err
		}

		fmt.Println("Pessoa escaneada:", pessoa.Nome)
		// Converter NullTime para *time.Time
		if deletadoEm.Valid {
			pessoa.Deletado_em = &deletadoEm.Time
		}
		if atualizadoEm.Valid {
			pessoa.Atualizado_em = &atualizadoEm.Time
		}

		// Buscar telefones da pessoa
		telefones, err := r.getTelefonesByPessoaId(db, pessoa.Id)
		if err != nil {
			return nil, err
		}
		pessoa.Telefones = telefones

		pessoas = append(pessoas, pessoa)
	}

	fmt.Println("Total de pessoas encontradas:", len(pessoas))
	// Verificar se houve erro durante a iteração
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pessoas, nil
}

// GetPessoaById retorna uma pessoa específica por ID
func (r *GetPessoasRepository) GetPessoaById(id int) (*models.Pessoa, error) {
	db := database.ConectarDB()
	defer db.Close()

	query := `SELECT id, nome, descricao, ativo, altura_metros, nascimento, cep, 
			  deletado_em, atualizado_em FROM pessoas WHERE id = ? AND deletado_em IS NULL`

	var pessoa models.Pessoa
	var deletadoEm, atualizadoEm sql.NullTime

	err := db.QueryRow(query, id).Scan(
		&pessoa.Id,
		&pessoa.Nome,
		&pessoa.Descricao,
		&pessoa.Ativo,
		&pessoa.Altura_metros,
		&pessoa.Nascimento,
		&pessoa.Cep,
		&deletadoEm,
		&atualizadoEm,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("pessoa não encontrada")
		}
		return nil, err
	}

	// Converter NullTime para *time.Time
	if deletadoEm.Valid {
		pessoa.Deletado_em = &deletadoEm.Time
	}
	if atualizadoEm.Valid {
		pessoa.Atualizado_em = &atualizadoEm.Time
	}

	// Buscar telefones da pessoa
	telefones, err := r.getTelefonesByPessoaId(db, pessoa.Id)
	if err != nil {
		return nil, err
	}
	pessoa.Telefones = telefones

	return &pessoa, nil
}

// getTelefonesByPessoaId busca os telefones de uma pessoa específica
func (r *GetPessoasRepository) getTelefonesByPessoaId(db *sql.DB, pessoaId int) ([]models.Telefone, error) {
	query := `SELECT id, pessoa_id, telefone FROM telefones WHERE pessoa_id = ?`

	rows, err := db.Query(query, pessoaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var telefones []models.Telefone
	for rows.Next() {
		var telefone models.Telefone
		err := rows.Scan(&telefone.Id, &telefone.PessoaId, &telefone.Telefone)
		if err != nil {
			return nil, err
		}
		telefones = append(telefones, telefone)
	}

	// Verificar se houve erro durante a iteração
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return telefones, nil
}
