# LivrariaGo
Projeto de um CRUD de uma livraria para disciplina de Banco de Dados da UFPB

# Sistema de Gerenciamento de Livraria - Documentação e Requisitos

## 1. Documentação e Requisitos do Sistema CRUD

### 1.1 Visão Geral
Sistema para gerenciamento de estoque de livros e cadastro de clientes de uma livraria, permitindo operações completas de CRUD (Create, Read, Update, Delete) e geração de relatórios gerenciais.

### 1.2 Entidades Principais

#### 1.2.1 Livro (Estoque)
**Atributos:**
- `id`: Identificador único (int)
- `titulo`: Título do livro (string)
- `autor`: Nome do autor (string)
- `isbn`: Código ISBN (string)
- `preco`: Preço unitário (decimal)
- `quantidadeEstoque`: Quantidade disponível (int)
- `categoria`: Categoria/gênero (string)
- `editora`: Nome da editora (string)
- `anoPublicacao`: Ano de publicação (int)
- `dataUltimaAtualizacao`: Data da última modificação (datetime)

#### 1.2.2 Cliente
**Atributos:**
- `id`: Identificador único (int)
- `nome`: Nome completo (string)
- `email`: E-mail (string)
- `telefone`: Telefone de contato (string)
- `cpf`: CPF do cliente (string)
- `endereco`: Endereço completo (string)
- `dataNascimento`: Data de nascimento (datetime)
- `dataCadastro`: Data do cadastro (datetime)
- `ativo`: Status do cliente (boolean)

#### 1.2.3 Venda
**Atributos:**
- `id`: Identificador único (int)
- `clienteId`: ID do cliente (int)
- `dataVenda`: Data da venda (datetime)
- `valorTotal`: Valor total da venda (decimal)
- `formaPagamento`: Forma de pagamento (string)
- `status`: Status da venda (string)

#### 1.2.4 ItemVenda
**Atributos:**
- `id`: Identificador único (int)
- `vendaId`: ID da venda (int)
- `livroId`: ID do livro (int)
- `quantidade`: Quantidade vendida (int)
- `precoUnitario`: Preço unitário na venda (decimal)
- `subtotal`: Subtotal do item (decimal)

### 1.3 Operações CRUD Requeridas

#### Para cada entidade (Livro, Cliente, Venda):
1. **Inserir**: Adicionar novo registro
2. **Alterar**: Modificar registro existente
3. **Pesquisar por nome**: Buscar registros por nome/título
4. **Remover**: Excluir registro (lógica ou física)
5. **Listar todos**: Exibir todos os registros
6. **Exibir um**: Mostrar detalhes de um registro específico

### 1.4 Funcionalidades Adicionais
- Controle de estoque automático nas vendas
- Validações de dados de entrada
- Histórico de operações
- Sistema de backup
- Relatórios gerenciais

## 2. Modelagem UML das Classes

```
┌─────────────────────────────┐
│          Livro              │
├─────────────────────────────┤
│ - id: int                   │
│ - titulo: string            │
│ - autor: string             │
│ - isbn: string              │
│ - preco: decimal            │
│ - quantidadeEstoque: int    │
│ - categoria: string         │
│ - editora: string           │
│ - anoPublicacao: int        │
│ - dataUltimaAtualizacao     │
├─────────────────────────────┤
│ + inserir(): bool           │
│ + atualizar(): bool         │
│ + buscarPorTitulo(): List   │
│ + remover(): bool           │
│ + listarTodos(): List       │
│ + obterPorId(): Livro       │
│ + validarISBN(): bool       │
│ + atualizarEstoque(): bool  │
│ + calcularValorTotal(): dec │
└─────────────────────────────┘
              │
              │
┌─────────────────────────────┐
│         Cliente             │
├─────────────────────────────┤
│ - id: int                   │
│ - nome: string              │
│ - email: string             │
│ - telefone: string          │
│ - cpf: string               │
│ - endereco: string          │
│ - dataNascimento: datetime  │
│ - dataCadastro: datetime    │
│ - ativo: bool               │
├─────────────────────────────┤
│ + inserir(): bool           │
│ + atualizar(): bool         │
│ + buscarPorNome(): List     │
│ + remover(): bool           │
│ + listarTodos(): List       │
│ + obterPorId(): Cliente     │
│ + validarCPF(): bool        │
│ + validarEmail(): bool      │
│ + ativar(): bool            │
│ + desativar(): bool         │
└─────────────────────────────┘
              │
              │ 1
              │
              │ *
┌─────────────────────────────┐
│          Venda              │
├─────────────────────────────┤
│ - id: int                   │
│ - clienteId: int            │
│ - dataVenda: datetime       │
│ - valorTotal: decimal       │
│ - formaPagamento: string    │
│ - status: string            │
├─────────────────────────────┤
│ + inserir(): bool           │
│ + atualizar(): bool         │
│ + buscarPorData(): List     │
│ + remover(): bool           │
│ + listarTodos(): List       │
│ + obterPorId(): Venda       │
│ + calcularTotal(): decimal  │
│ + finalizarVenda(): bool    │
│ + cancelarVenda(): bool     │
└─────────────────────────────┘
              │ 1
              │
              │ *
┌─────────────────────────────┐
│        ItemVenda            │
├─────────────────────────────┤
│ - id: int                   │
│ - vendaId: int              │
│ - livroId: int              │
│ - quantidade: int           │
│ - precoUnitario: decimal    │
│ - subtotal: decimal         │
├─────────────────────────────┤
│ + inserir(): bool           │
│ + atualizar(): bool         │
│ + remover(): bool           │
│ + calcularSubtotal(): dec   │
│ + validarQuantidade(): bool │
└─────────────────────────────┘

┌─────────────────────────────┐
│      LivroManager           │
├─────────────────────────────┤
│ - livros: List<Livro>       │
├─────────────────────────────┤
│ + adicionarLivro(): bool    │
│ + editarLivro(): bool       │
│ + pesquisarPorTitulo(): List│
│ + removerLivro(): bool      │
│ + listarTodosLivros(): List │
│ + obterLivro(): Livro       │
│ + gerarRelatorioEstoque()   │
│ + verificarEstoqueBaixo()   │
│ + atualizarPrecos(): bool   │
└─────────────────────────────┘

┌─────────────────────────────┐
│      ClienteManager         │
├─────────────────────────────┤
│ - clientes: List<Cliente>   │
├─────────────────────────────┤
│ + adicionarCliente(): bool  │
│ + editarCliente(): bool     │
│ + pesquisarPorNome(): List  │
│ + removerCliente(): bool    │
│ + listarTodosClientes():List│
│ + obterCliente(): Cliente   │
│ + gerarRelatorioClientes()  │
│ + validarDadosCliente():bool│
│ + obterClientesAtivos():List│
└─────────────────────────────┘

┌─────────────────────────────┐
│       VendaManager          │
├─────────────────────────────┤
│ - vendas: List<Venda>       │
├─────────────────────────────┤
│ + registrarVenda(): bool    │
│ + editarVenda(): bool       │
│ + pesquisarVenda(): List    │
│ + cancelarVenda(): bool     │
│ + listarTodasVendas(): List │
│ + obterVenda(): Venda       │
│ + gerarRelatorioVendas()    │
│ + calcularFaturamento(): dec│
│ + obterVendasPorPeriodo()   │
└─────────────────────────────┘

┌─────────────────────────────┐
│     RelatorioManager        │
├─────────────────────────────┤
│ - livroManager: LivroMgr    │
│ - clienteManager: ClienteMgr│
│ - vendaManager: VendaMgr    │
├─────────────────────────────┤
│ + gerarRelatorioGeral()     │
│ + gerarRelatorioEstoque()   │
│ + gerarRelatorioVendas()    │
│ + gerarRelatorioClientes()  │
│ + gerarRelatorioFinanceiro()│
│ + exportarRelatorio(): bool │
│ + agendarRelatorio(): bool  │
└─────────────────────────────┘
```

## 3. Classes de Gerenciamento CRUD

### 3.1 LivroManager
Responsável por todas as operações relacionadas aos livros:
- Gerenciamento de estoque
- Controle de preços
- Relatórios de estoque
- Alertas de estoque baixo

### 3.2 ClienteManager
Responsável por todas as operações relacionadas aos clientes:
- Cadastro e validação de dados
- Controle de status (ativo/inativo)
- Histórico de compras
- Relatórios de clientes

### 3.3 VendaManager
Responsável por todas as operações relacionadas às vendas:
- Processamento de vendas
- Controle de estoque automático
- Cálculos financeiros
- Relatórios de vendas

### 3.4 RelatorioManager
Responsável pela geração de relatórios:
- Consolidação de dados
- Formatação de relatórios
- Exportação em diferentes formatos
- Agendamento de relatórios

## 4. Métodos Detalhados

### 4.1 Métodos da Classe Livro
```
+ validarDados(): bool
+ formatarISBN(): string
+ calcularDescontoCategoria(): decimal
+ verificarDisponibilidade(): bool
+ obterHistoricoPrecos(): List
+ atualizarDataModificacao(): void
+ compararPrecos(): int
+ clonar(): Livro
```

### 4.2 Métodos da Classe Cliente
```
+ calcularIdade(): int
+ formatarCPF(): string
+ validarDadosPessoais(): bool
+ obterHistoricoCompras(): List
+ calcularTicketMedio(): decimal
+ classificarCliente(): string
+ enviarNotificacao(): bool
+ atualizarUltimoContato(): void
```

### 4.3 Métodos da Classe Venda
```
+ adicionarItem(): bool
+ removerItem(): bool
+ aplicarDesconto(): void
+ calcularTroco(): decimal
+ validarFormaPagamento(): bool
+ gerarRecibo(): string
+ enviarComprovanteEmail(): bool
+ atualizarStatusPagamento(): void
```

## 5. Sistema de Relatórios

### 5.1 Relatório de Estoque
**Informações incluídas:**
- Total de livros cadastrados
- Valor total do estoque
- Livros com estoque baixo (< 5 unidades)
- Livros mais vendidos
- Livros sem movimento (últimos 30 dias)
- Distribuição por categoria
- Valor médio por livro
- Livros mais caros/baratos

### 5.2 Relatório de Vendas
**Informações incluídas:**
- Total de vendas no período
- Faturamento total
- Ticket médio
- Vendas por forma de pagamento
- Produtos mais vendidos
- Vendas por categoria
- Crescimento em relação ao período anterior
- Meta de vendas vs realizado

### 5.3 Relatório de Clientes
**Informações incluídas:**
- Total de clientes cadastrados
- Clientes ativos vs inativos
- Novos clientes no período
- Clientes por faixa etária
- Clientes por região
- Clientes mais compradores
- Ticket médio por cliente
- Clientes sem compras (últimos 90 dias)

### 5.4 Relatório Financeiro
**Informações incluídas:**
- Faturamento mensal/anual
- Margem de lucro por categoria
- Análise de rentabilidade
- Projeções de vendas
- Comparativo com períodos anteriores
- Indicadores de desempenho (KPIs)

## 6. Funcionalidades Extras dos Métodos

### 6.1 Validações
- Validação de CPF com dígito verificador
- Validação de ISBN com checksum
- Validação de e-mail com regex
- Validação de dados obrigatórios
- Validação de limites de estoque

### 6.2 Utilitários
- Formatação de valores monetários
- Formatação de datas
- Geração de códigos únicos
- Backup automático de dados
- Log de operações

### 6.3 Integração
- Exportação para Excel/PDF
- Importação de dados em lote
- API para integrações externas
- Sincronização com sistemas contábeis
- Notificações automáticas
