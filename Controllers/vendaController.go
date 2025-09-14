// Controllers/VendaController.go
package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"livrariago/Db"
	"livrariago/Models"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// função para ler todas as vendas
func ReadVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Db.Query(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas 
		ORDER BY data_venda DESC
	`)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Venda, 0)

	for rows.Next() {
		venda := Models.Venda{}
		err := rows.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
			&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, venda)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler uma venda por ID
func ReadByIdVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	rows, err := Db.Db.Query(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas WHERE id=$1
	`, id)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Venda, 0)

	for rows.Next() {
		venda := Models.Venda{}
		err := rows.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
			&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, venda)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler vendas por cliente
func ReadByClienteIdVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	clienteId := r.URL.Query().Get("cliente_id")

	rows, err := Db.Db.Query(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas WHERE cliente_id=$1 
		ORDER BY data_venda DESC
	`, clienteId)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Venda, 0)

	for rows.Next() {
		venda := Models.Venda{}
		err := rows.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
			&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, venda)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler vendas por status
func ReadByStatusVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	status := r.URL.Query().Get("status")

	// Validar status
	statusValidos := map[string]bool{
		"PENDENTE":   true,
		"CONFIRMADA": true,
		"CANCELADA":  true,
		"ESTORNADA":  true,
	}

	if !statusValidos[status] {
		http.Error(w, "Status inválido", http.StatusBadRequest)
		return
	}

	rows, err := Db.Db.Query(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas WHERE status=$1 
		ORDER BY data_venda DESC
	`, status)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Venda, 0)

	for rows.Next() {
		venda := Models.Venda{}
		err := rows.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
			&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, venda)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para ler vendas por período
func ReadByPeriodoVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	dataInicio := r.URL.Query().Get("data_inicio") // formato: 2024-01-01
	dataFim := r.URL.Query().Get("data_fim")       // formato: 2024-12-31

	if dataInicio == "" || dataFim == "" {
		http.Error(w, "data_inicio e data_fim são obrigatórios (formato: YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	rows, err := Db.Db.Query(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas 
		WHERE DATE(data_venda) BETWEEN $1 AND $2 
		ORDER BY data_venda DESC
	`, dataInicio, dataFim)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	data := make([]Models.Venda, 0)

	for rows.Next() {
		venda := Models.Venda{}
		err := rows.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
			&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)
		if err != nil {
			fmt.Println("Server failed to handler = ", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data = append(data, venda)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// função para criar uma nova venda
func CreateVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	venda := Models.Venda{}
	err := json.NewDecoder(r.Body).Decode(&venda)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validações básicas
	if venda.ClienteId <= 0 {
		http.Error(w, "cliente_id é obrigatório e deve ser maior que zero", http.StatusBadRequest)
		return
	}

	// Validar forma de pagamento
	formasValidas := map[string]bool{
		"DINHEIRO":       true,
		"CARTAO_CREDITO": true,
		"CARTAO_DEBITO":  true,
		"PIX":            true,
		"BOLETO":         true,
	}

	if !formasValidas[venda.FormaPagamento] {
		http.Error(w, "Forma de pagamento inválida", http.StatusBadRequest)
		return
	}

	// Definir valores padrão
	if venda.Status == "" {
		venda.Status = "PENDENTE"
	}
	if venda.DataVenda.IsZero() {
		venda.DataVenda = time.Now()
	}
	if venda.ValorTotal.IsZero() {
		venda.ValorTotal = decimal.NewFromInt(0)
	}

	// Verificar se cliente existe
	var clienteCount int
	err = Db.Db.QueryRow("SELECT COUNT(*) FROM clientes WHERE id = $1 AND ativo = true", venda.ClienteId).Scan(&clienteCount)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if clienteCount == 0 {
		http.Error(w, "Cliente não encontrado ou inativo", http.StatusBadRequest)
		return
	}

	_, err = Db.Db.Exec(`
		INSERT INTO vendas (cliente_id, data_venda, valor_total, forma_pagamento, status, observacoes) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		venda.ClienteId, venda.DataVenda, venda.ValorTotal, venda.FormaPagamento, venda.Status, venda.Observacoes)

	fmt.Println("Venda recebida: ", venda)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Erro ao inserir venda", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// função para atualizar uma venda
func UpdateVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Models.Venda{}

	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Buscar a venda atual
	row := Db.Db.QueryRow(`
		SELECT id, cliente_id, data_venda, valor_total, forma_pagamento, status, 
		       COALESCE(observacoes, '') as observacoes 
		FROM vendas WHERE id = $1
	`, id)

	venda := Models.Venda{}
	err = row.Scan(&venda.Id, &venda.ClienteId, &venda.DataVenda,
		&venda.ValorTotal, &venda.FormaPagamento, &venda.Status, &venda.Observacoes)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// Não permitir atualizar vendas confirmadas ou canceladas (regra de negócio)
	if venda.Status == "CONFIRMADA" {
		http.Error(w, "Não é possível alterar venda confirmada", http.StatusBadRequest)
		return
	}
	if venda.Status == "CANCELADA" || venda.Status == "ESTORNADA" {
		http.Error(w, "Não é possível alterar venda cancelada ou estornada", http.StatusBadRequest)
		return
	}

	// Atualizar campos que foram fornecidos
	if up.ClienteId > 0 {
		// Verificar se cliente existe
		var clienteCount int
		err = Db.Db.QueryRow("SELECT COUNT(*) FROM clientes WHERE id = $1 AND ativo = true", up.ClienteId).Scan(&clienteCount)
		if err != nil || clienteCount == 0 {
			http.Error(w, "Cliente não encontrado ou inativo", http.StatusBadRequest)
			return
		}
		venda.ClienteId = up.ClienteId
	}

	if !up.ValorTotal.IsZero() {
		venda.ValorTotal = up.ValorTotal
	}

	if up.FormaPagamento != "" {
		formasValidas := map[string]bool{
			"DINHEIRO":       true,
			"CARTAO_CREDITO": true,
			"CARTAO_DEBITO":  true,
			"PIX":            true,
			"BOLETO":         true,
		}
		if !formasValidas[up.FormaPagamento] {
			http.Error(w, "Forma de pagamento inválida", http.StatusBadRequest)
			return
		}
		venda.FormaPagamento = up.FormaPagamento
	}

	if up.Status != "" {
		statusValidos := map[string]bool{
			"PENDENTE":   true,
			"CONFIRMADA": true,
			"CANCELADA":  true,
			"ESTORNADA":  true,
		}
		if !statusValidos[up.Status] {
			http.Error(w, "Status inválido", http.StatusBadRequest)
			return
		}
		venda.Status = up.Status
	}

	if up.Observacoes != "" {
		venda.Observacoes = up.Observacoes
	}

	_, err = Db.Db.Exec(`
		UPDATE vendas 
		SET cliente_id = $1, valor_total = $2, forma_pagamento = $3, status = $4, observacoes = $5 
		WHERE id = $6`,
		venda.ClienteId, venda.ValorTotal, venda.FormaPagamento, venda.Status, venda.Observacoes, venda.Id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(venda)
}

// função para confirmar uma venda (alterar status para CONFIRMADA)
func ConfirmarVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := Db.Db.Exec(`
		UPDATE vendas 
		SET status = 'CONFIRMADA' 
		WHERE id = $1 AND status = 'PENDENTE'`,
		id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Venda confirmada com sucesso"}`)
}

// função para cancelar uma venda
func CancelarVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	motivo := r.URL.Query().Get("motivo")

	// Atualizar observações com o motivo do cancelamento
	observacoes := "CANCELADA"
	if motivo != "" {
		observacoes += " - Motivo: " + motivo
	}

	_, err := Db.Db.Exec(`
		UPDATE vendas 
		SET status = 'CANCELADA', observacoes = COALESCE(observacoes, '') || ' | ' || $1 
		WHERE id = $2 AND status IN ('PENDENTE', 'CONFIRMADA')`,
		observacoes, id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Venda cancelada com sucesso"}`)
}

// função para deletar uma venda (só permite se status for PENDENTE)
func DeleteVenda(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	// Verificar se a venda pode ser deletada
	var status string
	err := Db.Db.QueryRow("SELECT status FROM vendas WHERE id = $1", id).Scan(&status)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if status != "PENDENTE" {
		http.Error(w, "Só é possível deletar vendas com status PENDENTE", http.StatusBadRequest)
		return
	}

	_, err = Db.Db.Exec("DELETE FROM vendas WHERE id = $1", id)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// função para obter relatório de vendas resumido
func RelatorioVendas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	dataInicio := r.URL.Query().Get("data_inicio")
	dataFim := r.URL.Query().Get("data_fim")

	var whereClause string
	var args []interface{}

	if dataInicio != "" && dataFim != "" {
		whereClause = "WHERE DATE(data_venda) BETWEEN $1 AND $2"
		args = append(args, dataInicio, dataFim)
	}

	query := fmt.Sprintf(`
		SELECT 
			COUNT(*) as total_vendas,
			COUNT(CASE WHEN status = 'CONFIRMADA' THEN 1 END) as vendas_confirmadas,
			COUNT(CASE WHEN status = 'PENDENTE' THEN 1 END) as vendas_pendentes,
			COUNT(CASE WHEN status = 'CANCELADA' THEN 1 END) as vendas_canceladas,
			COALESCE(SUM(CASE WHEN status = 'CONFIRMADA' THEN valor_total END), 0) as faturamento_total,
			COALESCE(AVG(CASE WHEN status = 'CONFIRMADA' THEN valor_total END), 0) as ticket_medio
		FROM vendas %s
	`, whereClause)

	var relatorio struct {
		TotalVendas       int             `json:"total_vendas"`
		VendasConfirmadas int             `json:"vendas_confirmadas"`
		VendasPendentes   int             `json:"vendas_pendentes"`
		VendasCanceladas  int             `json:"vendas_canceladas"`
		FaturamentoTotal  decimal.Decimal `json:"faturamento_total"`
		TicketMedio       decimal.Decimal `json:"ticket_medio"`
	}

	err := Db.Db.QueryRow(query, args...).Scan(
		&relatorio.TotalVendas, &relatorio.VendasConfirmadas,
		&relatorio.VendasPendentes, &relatorio.VendasCanceladas,
		&relatorio.FaturamentoTotal, &relatorio.TicketMedio)

	if err != nil {
		fmt.Println("Server failed to handler = ", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relatorio)
}
