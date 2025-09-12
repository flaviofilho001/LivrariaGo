package Models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Livro struct {
	id                int
	titulo            string
	autor             string
	isbn              string
	preco             decimal.Decimal
	quantidadeEstoque int
	categoria         string
	editora           string
	anoPublicacao     int

	dataUltimaAtualizacao time.Timer
}
