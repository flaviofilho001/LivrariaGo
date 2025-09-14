// Models/Livro.go
package Models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Livro struct {
	Id                    int             `json:"id"`
	Titulo                string          `json:"titulo"`
	Autor                 string          `json:"autor"`
	Isbn                  string          `json:"isbn"`
	Preco                 decimal.Decimal `json:"preco"`
	QuantidadeEstoque     int             `json:"quantidade_estoque"`
	Categoria             string          `json:"categoria"`
	Editora               string          `json:"editora"`
	AnoPublicacao         int             `json:"ano_publicacao"`
	DataUltimaAtualizacao time.Time       `json:"data_ultima_atualizacao"`
	Ativo                 bool            `json:"ativo"`
}
