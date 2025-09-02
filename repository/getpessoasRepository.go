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

// GetPessoas retorna todas as pessoas ativas - tenta procedure primeiro, depois query direta
func (r *GetPessoasRepository) GetPessoas() ([]models.Pessoa, error) {
	fmt.Println("=== INICIANDO BUSCA DE PESSOAS ===")

	db := database.ConectarDB()
	defer db.Close()

	fmt.Println("Conexão OK")

	// TENTATIVA 1: Usar stored procedure
	fmt.Println("Tentando com stored procedure...")
	pessoas, err := r.getPessoasWithProcedure(db)
	if err == nil {
		fmt.Println("Sucesso com stored procedure!")
		return pessoas, nil
	}

	fmt.Println("Falha na stored procedure:", err)
	fmt.Println("Tentando com query direta...")

	// TENTATIVA 2: Usar query direta (fallback)
	pessoas, err = r.getPessoasWithQuery(db)
	if err != nil {
		fmt.Println("Falha também na query direta:", err)
		return nil, err
	}

	fmt.Println("Sucesso com query direta!")
	return pessoas, nil
}

// GetPessoaById retorna uma pessoa específica por ID - tenta procedure primeiro, depois query direta
func (r *GetPessoasRepository) GetPessoaById(id int) (*models.Pessoa, error) {
	db := database.ConectarDB()
	defer db.Close()

	// TENTATIVA 1: Usar stored procedure
	fmt.Println("Tentando buscar pessoa por ID com stored procedure...")
	pessoa, err := r.getPessoaByIdWithProcedure(db, id)
	if err == nil {
		fmt.Println("Sucesso com stored procedure!")
		return pessoa, nil
	}

	fmt.Println("Falha na stored procedure:", err)
	fmt.Println("Tentando com query direta...")

	// TENTATIVA 2: Usar query direta (fallback)
	pessoa, err = r.getPessoaByIdWithQuery(db, id)
	if err != nil {
		fmt.Println("Falha também na query direta:", err)
		return nil, err
	}

	fmt.Println("Sucesso com query direta!")
	return pessoa, nil
}

// getPessoasWithProcedure - implementação usando stored procedure
func (r *GetPessoasRepository) getPessoasWithProcedure(db *sql.DB) ([]models.Pessoa, error) {
	query := `CALL sp_buscar_pessoas()`
	fmt.Println("Query (procedure):", query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro na procedure: %w", err)
	}
	defer rows.Close()

	return r.scanPessoas(db, rows)
}

// getPessoasWithQuery - implementação usando query direta
func (r *GetPessoasRepository) getPessoasWithQuery(db *sql.DB) ([]models.Pessoa, error) {
	query := `SELECT id, nome, descricao, ativo, altura_metros, nascimento, cep, 
			  deletado_em, atualizado_em FROM pessoas WHERE deletado_em IS NULL`
	fmt.Println("Query (direta):", query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro na query: %w", err)
	}
	defer rows.Close()

	return r.scanPessoas(db, rows)
}

// getPessoaByIdWithProcedure - busca por ID usando stored procedure
func (r *GetPessoasRepository) getPessoaByIdWithProcedure(db *sql.DB, id int) (*models.Pessoa, error) {
	query := `CALL sp_buscar_pessoa_por_id(?)`

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
		return nil, fmt.Errorf("erro na procedure: %w", err)
	}

	return r.buildPessoa(&pessoa, deletadoEm, atualizadoEm, db)
}

// getPessoaByIdWithQuery - busca por ID usando query direta
func (r *GetPessoasRepository) getPessoaByIdWithQuery(db *sql.DB, id int) (*models.Pessoa, error) {
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
		return nil, fmt.Errorf("erro na query: %w", err)
	}

	return r.buildPessoa(&pessoa, deletadoEm, atualizadoEm, db)
}

// scanPessoas - função auxiliar para escanear múltiplas pessoas (reutilizada)
func (r *GetPessoasRepository) scanPessoas(db *sql.DB, rows *sql.Rows) ([]models.Pessoa, error) {
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

		pessoaCompleta, err := r.buildPessoa(&pessoa, deletadoEm, atualizadoEm, db)
		if err != nil {
			return nil, err
		}

		pessoas = append(pessoas, *pessoaCompleta)
	}

	fmt.Println("Total de pessoas encontradas:", len(pessoas))

	// Verificar se houve erro durante a iteração
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pessoas, nil
}

// buildPessoa - função auxiliar para construir uma pessoa completa com telefones
func (r *GetPessoasRepository) buildPessoa(pessoa *models.Pessoa, deletadoEm, atualizadoEm sql.NullTime, db *sql.DB) (*models.Pessoa, error) {
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

	return pessoa, nil
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return telefones, nil
}
