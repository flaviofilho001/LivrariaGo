// Controllers/LivroController.go
package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"livrariago/Db"
	"livrariago/Models"
	"net/http"
	"strconv"
	"time"
)

// função para ler todos os livros
func ReadLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE ativo = true
	`)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler um livro por ID
func ReadByIdLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE id=$1
	`, id)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler livros por título
func ReadByTituloLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	titulo := r.URL.Query().Get("titulo")

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE titulo ILIKE $1 AND ativo = true
	`, "%"+titulo+"%")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler livros por autor
func ReadByAutorLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	autor := r.URL.Query().Get("autor")

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE autor ILIKE $1 AND ativo = true
	`, "%"+autor+"%")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler livros por categoria
func ReadByCategoriaLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	categoria := r.URL.Query().Get("categoria")

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE categoria ILIKE $1 AND ativo = true
	`, "%"+categoria+"%")
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler livros com estoque baixo
func ReadEstoqueBaixoLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	limite := r.URL.Query().Get("limite")
	if limite == "" {
		limite = "5" // Padrão: estoque baixo < 5
	}

	rows, err := Db.Db.Query(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE quantidade_estoque < $1 AND ativo = true
		ORDER BY quantidade_estoque ASC
	`, limite)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Livro, 0)

	for rows.Next() {
		livro := Models.Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
			&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
			&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para criar um novo livro
func CreateLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	livro := Models.Livro{}
	err := json.NewDecoder(r.Body).Decode(&livro)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validações básicas
	if livro.Titulo == "" || livro.Autor == "" || livro.Isbn == "" {
		http.Error(w, "Título, autor e ISBN são obrigatórios", http.StatusBadRequest)
		return
	}

	currentYear := time.Now().Year()
	if livro.AnoPublicacao < 1400 || livro.AnoPublicacao > currentYear {
		http.Error(w, "Ano de publicação inválido", http.StatusBadRequest)
		return
	}

	// Definir valores padrão
	if livro.QuantidadeEstoque < 0 {
		livro.QuantidadeEstoque = 0
	}
	livro.Ativo = true
	livro.DataUltimaAtualizacao = time.Now()

	_, err = Db.Db.Exec(`
		INSERT INTO livros (titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		                   editora, ano_publicacao, data_ultima_atualizacao, ativo) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		livro.Titulo, livro.Autor, livro.Isbn, livro.Preco, livro.QuantidadeEstoque,
		livro.Categoria, livro.Editora, livro.AnoPublicacao, livro.DataUltimaAtualizacao, livro.Ativo)

	fmt.Println("Livro recebido: ", livro)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		if err.Error() == `pq: duplicate key value violates unique constraint "livros_isbn_key"` {
			http.Error(w, "ISBN já cadastrado", http.StatusConflict)
		} else {
			http.Error(w, "Erro ao inserir livro", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// função para atualizar um livro
func UpdateLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Models.Livro{}

	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Buscar o livro atual
	row := Db.Db.QueryRow(`
		SELECT id, titulo, autor, isbn, preco, quantidade_estoque, categoria, 
		       editora, ano_publicacao, data_ultima_atualizacao, ativo 
		FROM livros WHERE id = $1
	`, id)

	livro := Models.Livro{}
	err = row.Scan(&livro.Id, &livro.Titulo, &livro.Autor, &livro.Isbn,
		&livro.Preco, &livro.QuantidadeEstoque, &livro.Categoria,
		&livro.Editora, &livro.AnoPublicacao, &livro.DataUltimaAtualizacao, &livro.Ativo)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// Atualizar campos que foram fornecidos
	if up.Titulo != "" {
		livro.Titulo = up.Titulo
	}
	if up.Autor != "" {
		livro.Autor = up.Autor
	}
	if up.Isbn != "" {
		livro.Isbn = up.Isbn
	}
	if !up.Preco.IsZero() {
		livro.Preco = up.Preco
	}
	if up.QuantidadeEstoque >= 0 {
		livro.QuantidadeEstoque = up.QuantidadeEstoque
	}
	if up.Categoria != "" {
		livro.Categoria = up.Categoria
	}
	if up.Editora != "" {
		livro.Editora = up.Editora
	}
	if up.AnoPublicacao > 0 {
		currentYear := time.Now().Year()
		if up.AnoPublicacao >= 1400 && up.AnoPublicacao <= currentYear {
			livro.AnoPublicacao = up.AnoPublicacao
		}
	}

	// Atualizar data de modificação
	livro.DataUltimaAtualizacao = time.Now()

	_, err = Db.Db.Exec(`
		UPDATE livros 
		SET titulo = $1, autor = $2, isbn = $3, preco = $4, quantidade_estoque = $5, 
		    categoria = $6, editora = $7, ano_publicacao = $8, data_ultima_atualizacao = $9 
		WHERE id = $10`,
		livro.Titulo, livro.Autor, livro.Isbn, livro.Preco, livro.QuantidadeEstoque,
		livro.Categoria, livro.Editora, livro.AnoPublicacao, livro.DataUltimaAtualizacao, livro.Id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		if err.Error() == `pq: duplicate key value violates unique constraint "livros_isbn_key"` {
			http.Error(w, "ISBN já cadastrado", http.StatusConflict)
		} else {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(livro)
}

// função para atualizar estoque de um livro
func UpdateEstoqueLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	quantidadeStr := r.URL.Query().Get("quantidade")

	quantidade, err := strconv.Atoi(quantidadeStr)
	if err != nil {
		http.Error(w, "Quantidade inválida", http.StatusBadRequest)
		return
	}

	_, err = Db.Db.Exec(`
		UPDATE livros 
		SET quantidade_estoque = $1, data_ultima_atualizacao = CURRENT_TIMESTAMP 
		WHERE id = $2 AND ativo = true`,
		quantidade, id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Estoque atualizado com sucesso"}`)
}

// função para deletar um livro (soft delete)
func DeleteLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	// Soft delete - apenas marca como inativo
	_, err := Db.Db.Exec(`
		UPDATE livros 
		SET ativo = false, data_ultima_atualizacao = CURRENT_TIMESTAMP 
		WHERE id = $1`,
		id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
