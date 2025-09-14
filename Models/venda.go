// Models/Venda.go
package Models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Venda struct {
	Id             int             `json:"id"`
	ClienteId      int             `json:"cliente_id"`
	DataVenda      time.Time       `json:"data_venda"`
	ValorTotal     decimal.Decimal `json:"valor_total"`
	FormaPagamento string          `json:"forma_pagamento"`
	Status         string          `json:"status"`
	Observacoes    string          `json:"observacoes"`
}
