package Models

import (
	"time"
)

type Cliente struct {
	Id             int
	Nome           string
	Email          string
	Telefone       string
	Cpf            string
	Endereco       string
	DataNascimento string
	DataCadastro   time.Time
	Ativo          bool
}
