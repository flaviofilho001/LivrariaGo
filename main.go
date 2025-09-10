package main

import (
	"fmt"
	"livrariago/Controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/clientes/update", Controllers.Update)
	http.HandleFunc("/clientes/delete", Controllers.Delete)
	http.HandleFunc("/clientes/read", Controllers.Read)
	http.HandleFunc("/clientes/create", Controllers.Create)
	http.HandleFunc("/clientes/readid", Controllers.ReadById)
	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)

}
