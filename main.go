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

	fmt.Println("Your connection in server 8080 was sucessfull")

	http.ListenAndServe(":8080", nil)

}
