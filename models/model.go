package models

type Pessoa struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Descricao     string `json:"descricao"`
	Deletado_em   string `json:"deletado_em"`
	Atualizado_em string `json:"atualizado_em"`
}
