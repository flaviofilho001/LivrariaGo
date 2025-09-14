// Controllers/ItemVendaController.go
package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"livrariago/Db"
	"livrariago/Models"
	"net/http"

	"github.com/shopspring/decimal"
)

// função para ler todos os itens de venda
func ReadItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Db.Query("SELECT id, venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto FROM itens_venda")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.ItemVenda, 0)

	for rows.Next() {
		item := Models.ItemVenda{}
		err := rows.Scan(&item.Id, &item.VendaId, &item.LivroId, &item.Quantidade,
			&item.PrecoUnitario, &item.Subtotal, &item.Desconto)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler um item de venda por ID
func ReadByIdItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	rows, err := Db.Db.Query("SELECT id, venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto FROM itens_venda WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.ItemVenda, 0)

	for rows.Next() {
		item := Models.ItemVenda{}
		err := rows.Scan(&item.Id, &item.VendaId, &item.LivroId, &item.Quantidade,
			&item.PrecoUnitario, &item.Subtotal, &item.Desconto)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler itens de venda por venda_id
func ReadByVendaIdItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vendaId := r.URL.Query().Get("venda_id")

	rows, err := Db.Db.Query("SELECT id, venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto FROM itens_venda WHERE venda_id=$1", vendaId)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.ItemVenda, 0)

	for rows.Next() {
		item := Models.ItemVenda{}
		err := rows.Scan(&item.Id, &item.VendaId, &item.LivroId, &item.Quantidade,
			&item.PrecoUnitario, &item.Subtotal, &item.Desconto)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler itens de venda por livro_id
func ReadByLivroIdItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	livroId := r.URL.Query().Get("livro_id")

	rows, err := Db.Db.Query("SELECT id, venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto FROM itens_venda WHERE livro_id=$1", livroId)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.ItemVenda, 0)

	for rows.Next() {
		item := Models.ItemVenda{}
		err := rows.Scan(&item.Id, &item.VendaId, &item.LivroId, &item.Quantidade,
			&item.PrecoUnitario, &item.Subtotal, &item.Desconto)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, item)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para criar um novo item de venda
func CreateItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	item := Models.ItemVenda{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validações básicas
	if item.VendaId <= 0 || item.LivroId <= 0 || item.Quantidade <= 0 || item.PrecoUnitario.LessThanOrEqual(decimal.Zero) {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Calcular subtotal se não foi fornecido
	if item.Subtotal.IsZero() {
		desconto := item.Desconto.Div(decimal.NewFromInt(100)) // Converter porcentagem para decimal
		item.Subtotal = item.PrecoUnitario.Mul(decimal.NewFromInt(int64(item.Quantidade))).Mul(decimal.NewFromInt(1).Sub(desconto))
	}

	_, err = Db.Db.Exec(`
		INSERT INTO itens_venda (venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		item.VendaId, item.LivroId, item.Quantidade, item.PrecoUnitario, item.Subtotal, item.Desconto)

	fmt.Println("Item de venda recebido: ", item)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Erro ao inserir item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// função para atualizar um item de venda
func UpdateItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Models.ItemVenda{}

	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Buscar o item atual
	row := Db.Db.QueryRow("SELECT id, venda_id, livro_id, quantidade, preco_unitario, subtotal, desconto FROM itens_venda WHERE id = $1", id)
	item := Models.ItemVenda{}
	err = row.Scan(&item.Id, &item.VendaId, &item.LivroId, &item.Quantidade,
		&item.PrecoUnitario, &item.Subtotal, &item.Desconto)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// Atualizar campos que foram fornecidos
	if up.VendaId > 0 {
		item.VendaId = up.VendaId
	}
	if up.LivroId > 0 {
		item.LivroId = up.LivroId
	}
	if up.Quantidade > 0 {
		item.Quantidade = up.Quantidade
	}
	if !up.PrecoUnitario.IsZero() {
		item.PrecoUnitario = up.PrecoUnitario
	}
	if !up.Desconto.IsZero() {
		item.Desconto = up.Desconto
	}

	// Recalcular subtotal
	desconto := item.Desconto.Div(decimal.NewFromInt(100))
	item.Subtotal = item.PrecoUnitario.Mul(decimal.NewFromInt(int64(item.Quantidade))).Mul(decimal.NewFromInt(1).Sub(desconto))

	_, err = Db.Db.Exec(`
		UPDATE itens_venda 
		SET venda_id = $1, livro_id = $2, quantidade = $3, preco_unitario = $4, subtotal = $5, desconto = $6 
		WHERE id = $7`,
		item.VendaId, item.LivroId, item.Quantidade, item.PrecoUnitario, item.Subtotal, item.Desconto, item.Id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// função para deletar um item de venda
func DeleteItemVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := Db.Db.Exec("DELETE FROM itens_venda WHERE id = $1", id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
