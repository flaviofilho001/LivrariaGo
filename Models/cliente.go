package Models

import (
	"time"
)

type Cliente struct {
	Id             int       `json:"id"`
	Nome           string    `json:"nome"`
	Email          string    `json:"email"`
	Telefone       string    `json:"telefone"`
	Cpf            string    `json:"cpf"`
	Endereco       string    `json:"endereco"`
	DataNascimento string    `json:"datanascimento"`
	DataCadastro   time.Time `json:"datacadastro"`
	Ativo          bool      `json:"ativo"`
}
