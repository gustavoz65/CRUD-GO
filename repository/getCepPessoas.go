package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type GetCepPessoasRepository struct {
}

type EnderecoViaCEP struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Erro       bool   `json:"erro"`
}

func (r *GetCepPessoasRepository) GetCepPessoas(cep string) (*EnderecoViaCEP, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endereco EnderecoViaCEP
	if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
		return nil, err
	}

	if endereco.Erro {
		return nil, fmt.Errorf("Cep não encontrado")
	}

	return &endereco, nil
}
