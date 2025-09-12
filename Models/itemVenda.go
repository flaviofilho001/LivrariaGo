package Models

import (
	"github.com/shopspring/decimal"
)

type ItemVenda struct {
	id            int
	vendaId       int
	livroId       int
	quantidade    int
	precoUnitario decimal.Decimal
	subtotal      decimal.Decimal
}
