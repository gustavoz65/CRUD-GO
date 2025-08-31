package models

import "time"

type Pessoa struct {
	Id            int        `json:"id"`
	Nome          string     `json:"nome"`
	Descricao     string     `json:"descricao"`
	Deletado_em   *time.Time `json:"deletado_em"`
	Atualizado_em *time.Time `json:"atualizado_em"`
	Ativo         bool       `json:"ativo"`
	Altura_metros float64    `json:"altura_metros"`
	Nascimento    time.Time  `json:"nascimento"`
	Cep           string     `json:"cep"`
	Telefones     []Telefone `json:"telefones"`
}

type Telefone struct {
	Id        int       `json:"id"`
	PessoaId  int       `json:"pessoa_id"`
	Telefone  string    `json:"telefone"`
	Criado_em time.Time `json:"criado_em"`
}
