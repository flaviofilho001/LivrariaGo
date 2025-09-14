CREATE TABLE IF NOT EXISTS clientes (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  telefone VARCHAR(20) NOT NULL,
  cpf VARCHAR(14) UNIQUE NOT NULL,
  endereco VARCHAR(200) NOT NULL,
  dataNascimento DATE NOT NULL,
  dataCadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ativo BOOLEAN DEFAULT TRUE
);

-- Tabela de Livros (Estoque)
CREATE TABLE IF NOT EXISTS livros (
  id SERIAL PRIMARY KEY,
  titulo VARCHAR(200) NOT NULL,
  autor VARCHAR(100) NOT NULL,
  isbn VARCHAR(20) UNIQUE NOT NULL,
  preco DECIMAL(10,2) NOT NULL CHECK (preco > 0),
  quantidade_estoque INTEGER NOT NULL DEFAULT 0 CHECK (quantidade_estoque >= 0),
  categoria VARCHAR(50) NOT NULL,
  editora VARCHAR(100) NOT NULL,
  ano_publicacao INTEGER NOT NULL CHECK (
    ano_publicacao > 1400 
    AND ano_publicacao <= EXTRACT(YEAR FROM CURRENT_DATE)
  ),
  data_ultima_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ativo BOOLEAN DEFAULT TRUE
);

-- Tabela de Vendas
CREATE TABLE IF NOT EXISTS vendas (
  id SERIAL PRIMARY KEY,
  cliente_id INTEGER NOT NULL,
  data_venda TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  valor_total DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (valor_total >= 0),
  forma_pagamento VARCHAR(20) NOT NULL CHECK (
    forma_pagamento IN ('DINHEIRO', 'CARTAO_CREDITO', 'CARTAO_DEBITO', 'PIX', 'BOLETO')
  ),
  status VARCHAR(20) NOT NULL DEFAULT 'PENDENTE' CHECK (
    status IN ('PENDENTE', 'CONFIRMADA', 'CANCELADA', 'ESTORNADA')
  ),
  observacoes TEXT,
  FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE RESTRICT
);

-- Tabela de Itens de Venda
CREATE TABLE IF NOT EXISTS itens_venda (
  id SERIAL PRIMARY KEY,
  venda_id INTEGER NOT NULL,
  livro_id INTEGER NOT NULL,
  quantidade INTEGER NOT NULL CHECK (quantidade > 0),
  preco_unitario DECIMAL(10,2) NOT NULL CHECK (preco_unitario > 0),
  subtotal DECIMAL(10,2) NOT NULL CHECK (subtotal >= 0),
  desconto DECIMAL(5,2) DEFAULT 0 CHECK (desconto >= 0 AND desconto <= 100),
  FOREIGN KEY (venda_id) REFERENCES vendas(id) ON DELETE CASCADE,
  FOREIGN KEY (livro_id) REFERENCES livros(id) ON DELETE RESTRICT
);

-- Tabela de Categorias
CREATE TABLE IF NOT EXISTS categorias (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL
);
