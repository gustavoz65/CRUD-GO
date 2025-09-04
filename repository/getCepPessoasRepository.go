package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gustavoz65/CRUD-NBS-GO/database"
)

type GetCepPessoasRepository struct {
}

type EnderecoViaCEP struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Erro       string `json:"erro"`
}

// GetCepPessoasById busca o CEP de uma pessoa pelo ID e consulta o ViaCEP
func (r *GetCepPessoasRepository) GetCepPessoasById(pessoaId int) (*EnderecoViaCEP, error) {
	// 1. Conecta ao banco de dados
	db := database.ConectarDB()
	defer db.Close()

	// 2. Busca a pessoa pelo ID para obter o CEP
	query := `SELECT cep FROM pessoas WHERE id = ? AND deletado_em IS NULL`
	var cep string

	// 3. Executa a query e escaneia o resultado
	err := db.QueryRow(query, pessoaId).Scan(&cep)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("pessoa não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar pessoa: %w", err)
	}

	// 4. Verifica se o CEP não está vazio
	if cep == "" {
		return nil, errors.New("pessoa não possui CEP cadastrado")
	}

	// 5. Consulta o ViaCEP com o CEP encontrado
	return r.consultarViaCEP(cep)
}

// consultarViaCEP faz a consulta na API do ViaCEP
func (r *GetCepPessoasRepository) consultarViaCEP(cep string) (*EnderecoViaCEP, error) {
	// 6. Monta a URL da API do ViaCEP
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	// 7. Faz a requisição HTTP GET para o ViaCEP
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar ViaCEP: %w", err)
	}
	defer resp.Body.Close()

	// 8. Decodifica a resposta JSON do ViaCEP
	var endereco EnderecoViaCEP
	if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta do ViaCEP: %w", err)
	}

	// 9. Verifica se o CEP foi encontrado no ViaCEP
	if endereco.Erro == "true" {
		return nil, fmt.Errorf("CEP %s não encontrado no ViaCEP", cep)
	}

	// 10. Retorna o endereço encontrado
	return &endereco, nil
}
