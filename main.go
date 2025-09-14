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
	http.HandleFunc("/clientes/update", Controllers.Update)
	http.HandleFunc("/clientes/delete", Controllers.Delete)
	http.HandleFunc("/clientes/read", Controllers.Read)
	http.HandleFunc("/clientes/create", Controllers.Create)
	http.HandleFunc("/clientes/readid", Controllers.ReadById)
	http.HandleFunc("/clientes/readnome", Controllers.ReadByNome)

	http.HandleFunc("/categorias/update", Controllers.UpdateCategoria)
	http.HandleFunc("/categorias/delete", Controllers.DeleteCategoria)
	http.HandleFunc("/categorias/read", Controllers.ReadCategoria)
	http.HandleFunc("/categorias/create", Controllers.CreateCategoria)
	http.HandleFunc("/categorias/readid", Controllers.ReadByIdCategoria)
	http.HandleFunc("/categorias/readnome", Controllers.ReadByNomeCategoria)

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)

}
