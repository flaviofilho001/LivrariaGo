package Models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Venda struct {
	id             int
	clienteId      int
	dataVenda      time.Time
	valorTotal     decimal.Decimal
	formaPagamento string
	status         string
}
