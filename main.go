package main

import (
	"fmt"
	"livrariago/Controllers"
	"net/http"
)

func main() {

	//da um cd LivrariaGo, ou cd c:local que o arquivo tá
	//pra iniciar tu faz docker-compose up --build
	//Quando for chamar as funções no postman ou no goole usa localhost:8080/clientes/aqui a função que tu quer
	// Caso seja alguma que precise do id, tu passa assim localhost:8080/clientes/readid?id=1

	// ROTAS PARA CLIENTES
	http.HandleFunc("/clientes/create", Controllers.Create)
	http.HandleFunc("/clientes/read", Controllers.Read)
	http.HandleFunc("/clientes/readid", Controllers.ReadById)
	http.HandleFunc("/clientes/readnome", Controllers.ReadByNome)
	http.HandleFunc("/clientes/update", Controllers.Update)
	http.HandleFunc("/clientes/delete", Controllers.Delete)

	// ROTAS PARA CATEGORIAS
	http.HandleFunc("/categorias/create", Controllers.CreateCategoria)
	http.HandleFunc("/categorias/read", Controllers.ReadCategoria)
	http.HandleFunc("/categorias/readid", Controllers.ReadByIdCategoria)
	http.HandleFunc("/categorias/readnome", Controllers.ReadByNomeCategoria)
	http.HandleFunc("/categorias/update", Controllers.UpdateCategoria)
	http.HandleFunc("/categorias/delete", Controllers.DeleteCategoria)

	// ROTAS PARA LIVROS
	http.HandleFunc("/livros/create", Controllers.CreateLivro)
	http.HandleFunc("/livros/read", Controllers.ReadLivro)
	http.HandleFunc("/livros/readid", Controllers.ReadByIdLivro)
	http.HandleFunc("/livros/readtitulo", Controllers.ReadByTituloLivro)       // ?titulo=nome_do_livro
	http.HandleFunc("/livros/readautor", Controllers.ReadByAutorLivro)         // ?autor=nome_do_autor
	http.HandleFunc("/livros/readcategoria", Controllers.ReadByCategoriaLivro) // ?categoria=nome_categoria
	http.HandleFunc("/livros/estoquebaixo", Controllers.ReadEstoqueBaixoLivro) // ?limite=5
	http.HandleFunc("/livros/update", Controllers.UpdateLivro)
	http.HandleFunc("/livros/updateestoque", Controllers.UpdateEstoqueLivro) // ?id=1&quantidade=100
	http.HandleFunc("/livros/delete", Controllers.DeleteLivro)

	// ROTAS PARA VENDAS
	http.HandleFunc("/vendas/create", Controllers.CreateVenda)
	http.HandleFunc("/vendas/read", Controllers.ReadVenda)
	http.HandleFunc("/vendas/readid", Controllers.ReadByIdVenda)
	http.HandleFunc("/vendas/readcliente", Controllers.ReadByClienteIdVenda) // ?cliente_id=1
	http.HandleFunc("/vendas/readstatus", Controllers.ReadByStatusVenda)     // ?status=CONFIRMADA
	http.HandleFunc("/vendas/readperiodo", Controllers.ReadByPeriodoVenda)   // ?data_inicio=2024-01-01&data_fim=2024-12-31
	http.HandleFunc("/vendas/update", Controllers.UpdateVenda)
	http.HandleFunc("/vendas/confirmar", Controllers.ConfirmarVenda) // ?id=1
	http.HandleFunc("/vendas/cancelar", Controllers.CancelarVenda)   // ?id=1&motivo=motivo_cancelamento
	http.HandleFunc("/vendas/delete", Controllers.DeleteVenda)
	http.HandleFunc("/vendas/relatorio", Controllers.RelatorioVendas) // ?data_inicio=2024-01-01&data_fim=2024-12-31

	// ROTAS PARA ITENS DE VENDA
	http.HandleFunc("/itensvenda/create", Controllers.CreateItemVenda)
	http.HandleFunc("/itensvenda/read", Controllers.ReadItemVenda)
	http.HandleFunc("/itensvenda/readid", Controllers.ReadByIdItemVenda)
	http.HandleFunc("/itensvenda/readvenda", Controllers.ReadByVendaIdItemVenda) // ?venda_id=1
	http.HandleFunc("/itensvenda/readlivro", Controllers.ReadByLivroIdItemVenda) // ?livro_id=1
	http.HandleFunc("/itensvenda/update", Controllers.UpdateItemVenda)
	http.HandleFunc("/itensvenda/delete", Controllers.DeleteItemVenda)

	fmt.Println("===========================================")
	fmt.Println("🚀 SERVIDOR LIVRARIA RODANDO NA PORTA 8080")
	fmt.Println("===========================================")
	fmt.Println("")
	fmt.Println("📚 ENDPOINTS DISPONÍVEIS:")
	fmt.Println("")
	fmt.Println("👥 CLIENTES:")
	fmt.Println("   GET    /clientes/read")
	fmt.Println("   GET    /clientes/readid?id=1")
	fmt.Println("   GET    /clientes/readnome?nome=João")
	fmt.Println("   POST   /clientes/create")
	fmt.Println("   PUT    /clientes/update?id=1")
	fmt.Println("   DELETE /clientes/delete?id=1")
	fmt.Println("")
	fmt.Println("🏷️  CATEGORIAS:")
	fmt.Println("   GET    /categorias/read")
	fmt.Println("   GET    /categorias/readid?id=1")
	fmt.Println("   GET    /categorias/readnome?nome=Literatura")
	fmt.Println("   POST   /categorias/create")
	fmt.Println("   PUT    /categorias/update?id=1")
	fmt.Println("   DELETE /categorias/delete?id=1")
	fmt.Println("")
	fmt.Println("📖 LIVROS:")
	fmt.Println("   GET    /livros/read")
	fmt.Println("   GET    /livros/readid?id=1")
	fmt.Println("   GET    /livros/readtitulo?titulo=Dom%20Casmurro")
	fmt.Println("   GET    /livros/readautor?autor=Machado")
	fmt.Println("   GET    /livros/readcategoria?categoria=Literatura")
	fmt.Println("   GET    /livros/estoquebaixo?limite=5")
	fmt.Println("   POST   /livros/create")
	fmt.Println("   PUT    /livros/update?id=1")
	fmt.Println("   PATCH  /livros/updateestoque?id=1&quantidade=100")
	fmt.Println("   DELETE /livros/delete?id=1")
	fmt.Println("")
	fmt.Println("💰 VENDAS:")
	fmt.Println("   GET    /vendas/read")
	fmt.Println("   GET    /vendas/readid?id=1")
	fmt.Println("   GET    /vendas/readcliente?cliente_id=1")
	fmt.Println("   GET    /vendas/readstatus?status=CONFIRMADA")
	fmt.Println("   GET    /vendas/readperiodo?data_inicio=2024-01-01&data_fim=2024-12-31")
	fmt.Println("   GET    /vendas/relatorio?data_inicio=2024-01-01&data_fim=2024-12-31")
	fmt.Println("   POST   /vendas/create")
	fmt.Println("   PUT    /vendas/update?id=1")
	fmt.Println("   PATCH  /vendas/confirmar?id=1")
	fmt.Println("   PATCH  /vendas/cancelar?id=1&motivo=Cliente%20desistiu")
	fmt.Println("   DELETE /vendas/delete?id=1")
	fmt.Println("")
	fmt.Println("📋 ITENS DE VENDA:")
	fmt.Println("   GET    /itensvenda/read")
	fmt.Println("   GET    /itensvenda/readid?id=1")
	fmt.Println("   GET    /itensvenda/readvenda?venda_id=1")
	fmt.Println("   GET    /itensvenda/readlivro?livro_id=1")
	fmt.Println("   POST   /itensvenda/create")
	fmt.Println("   PUT    /itensvenda/update?id=1")
	fmt.Println("   DELETE /itensvenda/delete?id=1")
	fmt.Println("")
	fmt.Println("===========================================")
	fmt.Println("🔗 Acesse: http://localhost:8080")
	fmt.Println("===========================================")

	http.ListenAndServe(":8080", nil)

}
