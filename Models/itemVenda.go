// Models/ItemVenda.go
package Models

import (
	"github.com/shopspring/decimal"
)

type ItemVenda struct {
	Id            int             `json:"id"`
	VendaId       int             `json:"venda_id"`
	LivroId       int             `json:"livro_id"`
	Quantidade    int             `json:"quantidade"`
	PrecoUnitario decimal.Decimal `json:"preco_unitario"`
	Subtotal      decimal.Decimal `json:"subtotal"`
	Desconto      decimal.Decimal `json:"desconto"`
}
