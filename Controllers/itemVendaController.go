package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"livrariago/Db"
	"livrariago/Models"
	"net/http"
	"time"
)

// função para ler as pessoas na tabela
func ReadItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Db.Query("SELECT * FROM clientes")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Cliente, 0)

	for rows.Next() {
		clientes := Models.Cliente{}
		err := rows.Scan(&clientes.Id, &clientes.Nome, &clientes.Email, &clientes.Telefone, &clientes.Cpf, &clientes.Endereco, &clientes.DataNascimento, &clientes.DataCadastro, &clientes.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, clientes)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler uma pessoa na tabela
func ReadByIdItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	rows, err := Db.Db.Query("SELECT * FROM clientes WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Cliente, 0)

	for rows.Next() {
		clientes := Models.Cliente{}
		err := rows.Scan(&clientes.Id, &clientes.Nome, &clientes.Email, &clientes.Telefone, &clientes.Cpf, &clientes.Endereco, &clientes.DataNascimento, &clientes.DataCadastro, &clientes.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, clientes)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ReadByNomeItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	nome := r.URL.Query().Get("nome")

	rows, err := Db.Db.Query("SELECT * FROM clientes WHERE nome=$1", nome)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	defer rows.Close()

	data := make([]Models.Cliente, 0)

	for rows.Next() {
		clientes := Models.Cliente{}
		err := rows.Scan(&clientes.Id, &clientes.Nome, &clientes.Email, &clientes.Telefone, &clientes.Cpf, &clientes.Endereco, &clientes.DataNascimento, &clientes.DataCadastro, &clientes.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			return
		}

		data = append(data, clientes)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func CreateItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	c := Models.Cliente{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}
	c.DataCadastro = time.Now()
	t, err := time.Parse("2006-01-02", c.DataNascimento)
	if err != nil {
		fmt.Println("Erro ao converter data:", err)
		return
	}

	_, err = Db.Db.Exec("INSERT INTO clientes (nome,email,telefone,cpf,endereco,dataNascimento,dataCadastro,ativo) VALUES ($1, $2, $3,$4,$5,$6,$7,$8)", c.Nome, c.Email, c.Telefone, c.Cpf, c.Endereco, t, c.DataCadastro, c.Ativo)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// função para atualizar uma pessoa na tabela
func UpdateItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Models.Cliente{}

	err := json.NewDecoder(r.Body).Decode(&up)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	row := Db.Db.QueryRow("SELECT * FROM clientes WHERE id = $1", id)
	c := Models.Cliente{}
	err = row.Scan(&c.Id, &c.Nome, &c.Email, &c.Telefone, &c.Cpf, &c.Endereco, &c.DataNascimento, &c.DataCadastro, &c.Ativo)

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
	if up.Email != "" {
		c.Email = up.Email
	}
	if up.Telefone != "" {
		c.Telefone = up.Telefone
	}
	if up.Cpf != "" {
		c.Cpf = up.Cpf
	}
	if up.Endereco != "" {
		c.Endereco = up.Endereco
	}
	if up.DataNascimento != "" { // só altera se não for "zero date"
		c.DataNascimento = up.DataNascimento
	}
	if up.Ativo != c.Ativo { // bool não tem "vazio", então compara
		c.Ativo = up.Ativo
	}

	_, err = Db.Db.Exec("UPDATE clientes SET nome = $1, email = $2,telefone = $3,cpf = $4 ,endereco = $5,dataNascimento = $6 ,dataCadastro = $7,ativo = $8 WHERE id=$9", c.Nome, c.Email, c.Telefone, c.Cpf, c.Endereco, c.DataNascimento, c.DataCadastro, c.Ativo, c.Id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

// função para detetar uma pessoa na tabela
func DeleteItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := Db.Db.Exec("DELETE FROM clientes WHERE id = $1 ", id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
