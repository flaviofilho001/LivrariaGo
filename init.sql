-- Limpar tabelas existentes se necessário (descomente se precisar)
-- DROP TABLE IF EXISTS itens_venda CASCADE;
-- DROP TABLE IF EXISTS vendas CASCADE;
-- DROP TABLE IF EXISTS livros CASCADE;
-- DROP TABLE IF EXISTS clientes CASCADE;
-- DROP TABLE IF EXISTS categorias CASCADE;

-- Tabela de Clientes (corrigida para snake_case)
CREATE TABLE IF NOT EXISTS clientes (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  telefone VARCHAR(20) NOT NULL,
  cpf VARCHAR(14) UNIQUE NOT NULL,
  endereco VARCHAR(200) NOT NULL,
  data_nascimento DATE NOT NULL,
  data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ativo BOOLEAN DEFAULT TRUE
);

-- Tabela de Categorias
CREATE TABLE IF NOT EXISTS categorias (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL UNIQUE
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

-- ========================================
-- ÍNDICES PARA OTIMIZAÇÃO DE PERFORMANCE
-- ========================================

-- Índices para Clientes
CREATE INDEX IF NOT EXISTS idx_clientes_cpf ON clientes(cpf);
CREATE INDEX IF NOT EXISTS idx_clientes_email ON clientes(email);
CREATE INDEX IF NOT EXISTS idx_clientes_nome ON clientes(nome);
CREATE INDEX IF NOT EXISTS idx_clientes_ativo ON clientes(ativo);

-- Índices para Categorias
CREATE INDEX IF NOT EXISTS idx_categorias_nome ON categorias(nome);

-- Índices para Livros
CREATE INDEX IF NOT EXISTS idx_livros_isbn ON livros(isbn);
CREATE INDEX IF NOT EXISTS idx_livros_titulo ON livros(titulo);
CREATE INDEX IF NOT EXISTS idx_livros_autor ON livros(autor);
CREATE INDEX IF NOT EXISTS idx_livros_categoria ON livros(categoria);
CREATE INDEX IF NOT EXISTS idx_livros_ativo ON livros(ativo);
CREATE INDEX IF NOT EXISTS idx_livros_estoque ON livros(quantidade_estoque);

-- Índices para Vendas
CREATE INDEX IF NOT EXISTS idx_vendas_cliente_id ON vendas(cliente_id);
CREATE INDEX IF NOT EXISTS idx_vendas_data ON vendas(data_venda);
CREATE INDEX IF NOT EXISTS idx_vendas_status ON vendas(status);
CREATE INDEX IF NOT EXISTS idx_vendas_forma_pagamento ON vendas(forma_pagamento);

-- Índices para Itens de Venda
CREATE INDEX IF NOT EXISTS idx_itens_venda_venda_id ON itens_venda(venda_id);
CREATE INDEX IF NOT EXISTS idx_itens_venda_livro_id ON itens_venda(livro_id);

-- ========================================
-- DADOS DE EXEMPLO
-- ========================================

-- Inserir categorias de exemplo
INSERT INTO categorias (nome) VALUES 
('Literatura Brasileira'),
('Literatura Estrangeira'),
('Tecnologia'),
('Ficção Científica'),
('Romance'),
('Biografia'),
('História'),
('Filosofia'),
('Autoajuda'),
('Literatura Infantil')
ON CONFLICT (nome) DO NOTHING;

-- Inserir clientes de exemplo
INSERT INTO clientes (nome, email, telefone, cpf, endereco, data_nascimento) VALUES
('João Silva Santos', 'joao.silva@email.com', '(83) 99999-1111', '123.456.789-01', 'Rua das Flores, 123 - Centro, João Pessoa/PB', '1990-05-15'),
('Maria Santos Oliveira', 'maria.santos@email.com', '(83) 99999-2222', '987.654.321-09', 'Av. Principal, 456 - Bairro Novo, João Pessoa/PB', '1985-08-22'),
('Pedro Oliveira Costa', 'pedro.oliveira@email.com', '(83) 99999-3333', '111.222.333-44', 'Rua da Paz, 789 - Vila Esperança, João Pessoa/PB', '1992-12-03'),
('Ana Paula Ferreira', 'ana.paula@email.com', '(83) 99999-4444', '555.666.777-88', 'Rua do Sol, 321 - Tambaú, João Pessoa/PB', '1988-03-10'),
('Carlos Eduardo Lima', 'carlos.eduardo@email.com', '(83) 99999-5555', '999.888.777-66', 'Av. Epitácio Pessoa, 654 - Cabo Branco, João Pessoa/PB', '1995-11-28')
ON CONFLICT (cpf) DO NOTHING;

-- Inserir livros de exemplo
INSERT INTO livros (titulo, autor, isbn, preco, quantidade_estoque, categoria, editora, ano_publicacao) VALUES
('Dom Casmurro', 'Machado de Assis', '978-85-359-0277-5', 29.90, 50, 'Literatura Brasileira', 'Companhia das Letras', 1899),
('Clean Code: A Handbook of Agile Software Craftsmanship', 'Robert C. Martin', '978-0-13-235088-4', 89.90, 25, 'Tecnologia', 'Prentice Hall', 2008),
('O Pequeno Príncipe', 'Antoine de Saint-Exupéry', '978-85-359-0624-7', 24.90, 100, 'Literatura Infantil', 'Agir', 1943),
('Algoritmos e Estruturas de Dados em C', 'Thomas H. Cormen', '978-85-352-3699-6', 159.90, 15, 'Tecnologia', 'Campus', 2012),
('1984', 'George Orwell', '978-85-359-0336-9', 34.90, 75, 'Ficção Científica', 'Companhia das Letras', 1949),
('O Cortiço', 'Aluísio Azevedo', '978-85-359-0445-8', 27.50, 40, 'Literatura Brasileira', 'Ática', 1890),
('Sapiens: Uma Breve História da Humanidade', 'Yuval Noah Harari', '978-85-359-2213-1', 54.90, 30, 'História', 'Companhia das Letras', 2014),
('O Guia do Mochileiro das Galáxias', 'Douglas Adams', '978-85-359-0892-0', 32.90, 60, 'Ficção Científica', 'Arqueiro', 1979),
('Mindset: A Nova Psicologia do Sucesso', 'Carol S. Dweck', '978-85-7905-934-1', 39.90, 20, 'Autoajuda', 'Objetiva', 2017),
('Steve Jobs', 'Walter Isaacson', '978-85-359-1059-7', 49.90, 25, 'Biografia', 'Companhia das Letras', 2011)
ON CONFLICT (isbn) DO NOTHING;

-- Inserir vendas de exemplo
INSERT INTO vendas (cliente_id, forma_pagamento, status, observacoes) VALUES
(1, 'PIX', 'CONFIRMADA', 'Pagamento à vista'),
(2, 'CARTAO_CREDITO', 'CONFIRMADA', 'Parcelado em 3x'),
(3, 'DINHEIRO', 'PENDENTE', 'Aguardando confirmação'),
(1, 'CARTAO_DEBITO', 'CONFIRMADA', 'Compra presencial'),
(4, 'PIX', 'PENDENTE', 'Primeira compra do cliente')
ON CONFLICT DO NOTHING;

-- Inserir itens de venda de exemplo (valor_total será calculado automaticamente)
INSERT INTO itens_venda (venda_id, livro_id, quantidade, preco_unitario, desconto) VALUES
(1, 1, 2, 29.90, 0),    -- Dom Casmurro x2
(1, 3, 1, 24.90, 5),    -- O Pequeno Príncipe com 5% desconto
(2, 2, 1, 89.90, 0),    -- Clean Code
(2, 7, 1, 54.90, 10),   -- Sapiens com 10% desconto
(3, 5, 1, 34.90, 0),    -- 1984
(4, 6, 1, 27.50, 0),    -- O Cortiço
(4, 8, 2, 32.90, 0),    -- Guia do Mochileiro x2
(5, 9, 1, 39.90, 15)    -- Mindset com 15% desconto
ON CONFLICT DO NOTHING;



-- ========================================
-- COMENTÁRIOS FINAIS
-- ========================================

-- Para habilitar o controle automático de estoque, descomente o trigger acima
-- Para testar as funções, execute algumas operações de CRUD
-- As views criadas facilitam a geração de relatórios
-- Os índices otimizam as consultas mais frequentes

COMMENT ON TABLE clientes IS 'Tabela de cadastro de clientes da livraria';
COMMENT ON TABLE categorias IS 'Tabela de categorias de livros';
COMMENT ON TABLE livros IS 'Tabela de estoque de livros';
COMMENT ON TABLE vendas IS 'Tabela de registro de vendas';
COMMENT ON TABLE itens_venda IS 'Tabela de itens individuais de cada venda';

-- Finalizar transação
COMMIT;