package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"livrariago/Db"
	"livrariago/Models"
	"net/http"
)

// função para ler as categorias na tabela
func ReadCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Db.Query("SELECT * FROM categorias")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Categoria, 0)

	for rows.Next() {
		categorias := Models.Categoria{}
		err := rows.Scan(&categorias.Id, &categorias.Nome)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, categorias)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler uma categorias na tabela
func ReadByIdCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	rows, err := Db.Db.Query("SELECT * FROM categorias WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Categoria, 0)

	for rows.Next() {
		categorias := Models.Categoria{}
		err := rows.Scan(&categorias.Id, &categorias.Nome)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, categorias)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ReadByNomeCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	nome := r.URL.Query().Get("nome")

	rows, err := Db.Db.Query("SELECT * FROM categorias WHERE nome=$1", nome)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Categoria, 0)

	for rows.Next() {
		categorias := Models.Categoria{}
		err := rows.Scan(&categorias.Id, &categorias.Nome)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, categorias)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func CreateCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	c := Models.Categoria{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	_, err = Db.Db.Exec("INSERT INTO categorias (nome) VALUES ($1)", c.Nome)

	fmt.Println("Categoria recebida: ", c)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// função para atualizar uma categorias na tabela
func UpdateCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Models.Categoria{}

	err := json.NewDecoder(r.Body).Decode(&up)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	row := Db.Db.QueryRow("SELECT * FROM categorias WHERE id = $1", id)
	c := Models.Categoria{}
	err = row.Scan(&c.Id, &c.Nome)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if up.Nome != "" {
		c.Nome = up.Nome
	}

	_, err = Db.Db.Exec("UPDATE categorias SET nome = $1 WHERE id=$2", c.Nome, c.Id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

// função para detetar uma categorias na tabela
func DeleteCategoria(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := Db.Db.Exec("DELETE FROM categorias WHERE id = $1 ", id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
