-- ==========================================
-- SISTEMA DE LIVRARIA - BANCO DE DADOS
-- Versão: 2.0.0 (Compatível com Models Go)
-- Data: 2024-09-29
-- ==========================================

SET client_encoding = 'UTF8';
SET timezone = 'America/Sao_Paulo';

-- ==========================================
-- EXTENSÕES
-- ==========================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "unaccent";

-- ==========================================
-- LIMPEZA (se necessário)
-- ==========================================

DROP TABLE IF EXISTS itens_venda CASCADE;
DROP TABLE IF EXISTS vendas CASCADE;
DROP TABLE IF EXISTS livros CASCADE;
DROP TABLE IF EXISTS clientes CASCADE;
DROP TABLE IF EXISTS categorias CASCADE;
DROP TABLE IF EXISTS historico_estoque CASCADE;
DROP TABLE IF EXISTS log_operacoes CASCADE;

-- ==========================================
-- TABELAS PRINCIPAIS
-- ==========================================

-- Tabela Categorias (sem campo ativo)
CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL UNIQUE
);

-- Tabela Clientes
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    endereco VARCHAR(300) NOT NULL,
    datanascimento DATE NOT NULL,
    datacadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ativo BOOLEAN DEFAULT TRUE,
    
    -- Constraints
    CONSTRAINT chk_clientes_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT chk_clientes_cpf CHECK (LENGTH(REPLACE(REPLACE(REPLACE(cpf, '.', ''), '-', ''), ' ', '')) = 11),
    CONSTRAINT chk_clientes_datanascimento CHECK (datanascimento <= CURRENT_DATE - INTERVAL '16 years')
);

-- Tabela Livros
CREATE TABLE livros (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(300) NOT NULL,
    autor VARCHAR(200) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    preco DECIMAL(10,2) NOT NULL CHECK (preco > 0),
    quantidade_estoque INTEGER NOT NULL DEFAULT 0 CHECK (quantidade_estoque >= 0),
    categoria VARCHAR(100) NOT NULL,
    editora VARCHAR(150) NOT NULL,
    ano_publicacao INTEGER NOT NULL,
    data_ultima_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ativo BOOLEAN DEFAULT TRUE,
    
    -- Constraints
    CONSTRAINT chk_livros_ano_publicacao CHECK (
        ano_publicacao >= 1400 AND 
        ano_publicacao <= EXTRACT(YEAR FROM CURRENT_DATE)
    ),
    CONSTRAINT chk_livros_isbn CHECK (
        LENGTH(REPLACE(REPLACE(isbn, '-', ''), ' ', '')) IN (10, 13)
    )
);

-- Tabela Vendas
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER NOT NULL REFERENCES clientes(id) ON DELETE RESTRICT,
    data_venda TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valor_total DECIMAL(12,2) NOT NULL DEFAULT 0 CHECK (valor_total >= 0),
    forma_pagamento VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDENTE',
    observacoes TEXT,
    
    -- Constraints
    CONSTRAINT chk_vendas_forma_pagamento CHECK (
        forma_pagamento IN ('DINHEIRO', 'CARTAO_CREDITO', 'CARTAO_DEBITO', 'PIX', 'BOLETO', 'TRANSFERENCIA')
    ),
    CONSTRAINT chk_vendas_status CHECK (
        status IN ('PENDENTE', 'CONFIRMADA', 'CANCELADA', 'ESTORNADA', 'PROCESSANDO')
    )
);

-- Tabela Itens de Venda
CREATE TABLE itens_venda (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER NOT NULL REFERENCES vendas(id) ON DELETE CASCADE,
    livro_id INTEGER NOT NULL REFERENCES livros(id) ON DELETE RESTRICT,
    quantidade INTEGER NOT NULL CHECK (quantidade > 0),
    preco_unitario DECIMAL(10,2) NOT NULL CHECK (preco_unitario > 0),
    subtotal DECIMAL(12,2) NOT NULL CHECK (subtotal >= 0),
    desconto DECIMAL(10,2) DEFAULT 0 CHECK (desconto >= 0),
    
    -- Evitar duplicatas na mesma venda
    UNIQUE(venda_id, livro_id)
);

-- ==========================================
-- TABELAS AUXILIARES
-- ==========================================

-- Histórico de Estoque
CREATE TABLE historico_estoque (
    id SERIAL PRIMARY KEY,
    livro_id INTEGER NOT NULL REFERENCES livros(id),
    tipo_movimentacao VARCHAR(20) NOT NULL,
    quantidade_anterior INTEGER NOT NULL,
    quantidade_movimentada INTEGER NOT NULL,
    quantidade_atual INTEGER NOT NULL,
    motivo VARCHAR(200),
    usuario VARCHAR(100),
    data_movimentacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    venda_id INTEGER REFERENCES vendas(id),
    
    CONSTRAINT chk_tipo_movimentacao CHECK (
        tipo_movimentacao IN ('ENTRADA', 'SAIDA', 'AJUSTE', 'VENDA', 'ESTORNO', 'INVENTARIO')
    )
);

-- Log de Operações
CREATE TABLE log_operacoes (
    id SERIAL PRIMARY KEY,
    tabela VARCHAR(50) NOT NULL,
    operacao VARCHAR(10) NOT NULL,
    registro_id INTEGER NOT NULL,
    dados_anteriores JSONB,
    dados_novos JSONB,
    usuario VARCHAR(100),
    ip_address INET,
    data_operacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ==========================================
-- ÍNDICES
-- ==========================================

-- Clientes
CREATE INDEX idx_clientes_cpf ON clientes(cpf);
CREATE INDEX idx_clientes_email ON clientes(email);
CREATE INDEX idx_clientes_nome ON clientes USING gin(to_tsvector('portuguese', nome));
CREATE INDEX idx_clientes_ativo ON clientes(ativo);
CREATE INDEX idx_clientes_datacadastro ON clientes(datacadastro);

-- Categorias
CREATE INDEX idx_categorias_nome ON categorias(nome);

-- Livros
CREATE INDEX idx_livros_isbn ON livros(isbn);
CREATE INDEX idx_livros_titulo ON livros USING gin(to_tsvector('portuguese', titulo));
CREATE INDEX idx_livros_autor ON livros USING gin(to_tsvector('portuguese', autor));
CREATE INDEX idx_livros_categoria ON livros(categoria);
CREATE INDEX idx_livros_ativo ON livros(ativo);
CREATE INDEX idx_livros_estoque ON livros(quantidade_estoque);
CREATE INDEX idx_livros_preco ON livros(preco);

-- Vendas
CREATE INDEX idx_vendas_cliente_id ON vendas(cliente_id);
CREATE INDEX idx_vendas_data_venda ON vendas(data_venda);
CREATE INDEX idx_vendas_status ON vendas(status);
CREATE INDEX idx_vendas_forma_pagamento ON vendas(forma_pagamento);

-- Itens de Venda
CREATE INDEX idx_itens_venda_venda_id ON itens_venda(venda_id);
CREATE INDEX idx_itens_venda_livro_id ON itens_venda(livro_id);

-- Histórico de Estoque
CREATE INDEX idx_historico_estoque_livro_id ON historico_estoque(livro_id);
CREATE INDEX idx_historico_estoque_data ON historico_estoque(data_movimentacao);

-- ==========================================
-- FUNÇÕES
-- ==========================================

-- Função para validar CPF
CREATE OR REPLACE FUNCTION validar_cpf(cpf_input TEXT)
RETURNS BOOLEAN AS $$
DECLARE
    cpf_limpo TEXT;
    soma INTEGER := 0;
    resto INTEGER;
    digito1 INTEGER;
    digito2 INTEGER;
    i INTEGER;
BEGIN
    cpf_limpo := REGEXP_REPLACE(cpf_input, '\D', '', 'g');
    
    IF LENGTH(cpf_limpo) != 11 THEN
        RETURN FALSE;
    END IF;
    
    IF cpf_limpo ~ '^(.)\1{10}$' THEN
        RETURN FALSE;
    END IF;
    
    FOR i IN 1..9 LOOP
        soma := soma + CAST(SUBSTRING(cpf_limpo FROM i FOR 1) AS INTEGER) * (11 - i);
    END LOOP;
    
    resto := soma % 11;
    IF resto < 2 THEN
        digito1 := 0;
    ELSE
        digito1 := 11 - resto;
    END IF;
    
    IF digito1 != CAST(SUBSTRING(cpf_limpo FROM 10 FOR 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;
    
    soma := 0;
    FOR i IN 1..10 LOOP
        soma := soma + CAST(SUBSTRING(cpf_limpo FROM i FOR 1) AS INTEGER) * (12 - i);
    END LOOP;
    
    resto := soma % 11;
    IF resto < 2 THEN
        digito2 := 0;
    ELSE
        digito2 := 11 - resto;
    END IF;
    
    IF digito2 != CAST(SUBSTRING(cpf_limpo FROM 11 FOR 1) AS INTEGER) THEN
        RETURN FALSE;
    END IF;
    
    RETURN TRUE;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- Função para atualizar timestamp
CREATE OR REPLACE FUNCTION atualizar_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.data_ultima_atualizacao = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Função para calcular subtotal do item
CREATE OR REPLACE FUNCTION calcular_subtotal_item()
RETURNS TRIGGER AS $$
BEGIN
    -- Subtotal = (quantidade * preco_unitario) - desconto
    NEW.subtotal = (NEW.quantidade * NEW.preco_unitario) - NEW.desconto;
    
    IF NEW.subtotal < 0 THEN
        NEW.subtotal = 0;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Função para atualizar total da venda
CREATE OR REPLACE FUNCTION atualizar_total_venda()
RETURNS TRIGGER AS $$
DECLARE
    v_venda_id INTEGER;
    v_total DECIMAL(12,2);
BEGIN
    IF TG_OP = 'DELETE' THEN
        v_venda_id := OLD.venda_id;
    ELSE
        v_venda_id := NEW.venda_id;
    END IF;
    
    SELECT COALESCE(SUM(subtotal), 0) INTO v_total
    FROM itens_venda 
    WHERE venda_id = v_venda_id;
    
    UPDATE vendas 
    SET valor_total = v_total
    WHERE id = v_venda_id;
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

-- Função para controlar estoque nas vendas
CREATE OR REPLACE FUNCTION controlar_estoque_venda()
RETURNS TRIGGER AS $$
DECLARE
    v_status_anterior VARCHAR(20);
    v_status_novo VARCHAR(20);
BEGIN
    IF TG_OP = 'UPDATE' THEN
        v_status_anterior := OLD.status;
        v_status_novo := NEW.status;
        
        -- PENDENTE -> CONFIRMADA: diminui estoque
        IF v_status_anterior = 'PENDENTE' AND v_status_novo = 'CONFIRMADA' THEN
            UPDATE livros 
            SET quantidade_estoque = quantidade_estoque - iv.quantidade,
                data_ultima_atualizacao = CURRENT_TIMESTAMP
            FROM itens_venda iv 
            WHERE livros.id = iv.livro_id AND iv.venda_id = NEW.id;
            
            -- Registra no histórico
            INSERT INTO historico_estoque (livro_id, tipo_movimentacao, quantidade_anterior, 
                                         quantidade_movimentada, quantidade_atual, motivo, venda_id)
            SELECT l.id, 'VENDA', l.quantidade_estoque + iv.quantidade, -iv.quantidade, 
                   l.quantidade_estoque, 'Venda confirmada #' || NEW.id, NEW.id
            FROM livros l
            JOIN itens_venda iv ON l.id = iv.livro_id
            WHERE iv.venda_id = NEW.id;
            
        -- CONFIRMADA -> ESTORNADA: aumenta estoque
        ELSIF v_status_anterior = 'CONFIRMADA' AND v_status_novo = 'ESTORNADA' THEN
            UPDATE livros 
            SET quantidade_estoque = quantidade_estoque + iv.quantidade,
                data_ultima_atualizacao = CURRENT_TIMESTAMP
            FROM itens_venda iv 
            WHERE livros.id = iv.livro_id AND iv.venda_id = NEW.id;
            
            -- Registra no histórico
            INSERT INTO historico_estoque (livro_id, tipo_movimentacao, quantidade_anterior, 
                                         quantidade_movimentada, quantidade_atual, motivo, venda_id)
            SELECT l.id, 'ESTORNO', l.quantidade_estoque - iv.quantidade, iv.quantidade, 
                   l.quantidade_estoque, 'Estorno venda #' || NEW.id, NEW.id
            FROM livros l
            JOIN itens_venda iv ON l.id = iv.livro_id
            WHERE iv.venda_id = NEW.id;
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Função para histórico de estoque manual
CREATE OR REPLACE FUNCTION registrar_historico_estoque()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'UPDATE' AND OLD.quantidade_estoque != NEW.quantidade_estoque THEN
        INSERT INTO historico_estoque (
            livro_id, tipo_movimentacao, quantidade_anterior, 
            quantidade_movimentada, quantidade_atual, motivo
        ) VALUES (
            NEW.id, 
            'AJUSTE',
            OLD.quantidade_estoque,
            NEW.quantidade_estoque - OLD.quantidade_estoque,
            NEW.quantidade_estoque,
            'Ajuste manual de estoque'
        );
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Função de log
CREATE OR REPLACE FUNCTION log_operacao()
RETURNS TRIGGER AS $$
DECLARE
    v_dados_anteriores JSONB;
    v_dados_novos JSONB;
BEGIN
    IF TG_OP = 'DELETE' THEN
        v_dados_anteriores = row_to_json(OLD)::jsonb;
        v_dados_novos = NULL;
    ELSIF TG_OP = 'INSERT' THEN
        v_dados_anteriores = NULL;
        v_dados_novos = row_to_json(NEW)::jsonb;
    ELSIF TG_OP = 'UPDATE' THEN
        v_dados_anteriores = row_to_json(OLD)::jsonb;
        v_dados_novos = row_to_json(NEW)::jsonb;
    END IF;
    
    IF TG_TABLE_NAME IN ('vendas', 'clientes', 'livros') THEN
        INSERT INTO log_operacoes (
            tabela, operacao, registro_id, 
            dados_anteriores, dados_novos
        ) VALUES (
            TG_TABLE_NAME, TG_OP, 
            COALESCE(NEW.id, OLD.id),
            v_dados_anteriores, v_dados_novos
        );
    END IF;
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

-- ==========================================
-- TRIGGERS
-- ==========================================

-- Atualizar timestamp dos livros
CREATE TRIGGER tr_livros_atualizar_timestamp
    BEFORE UPDATE ON livros
    FOR EACH ROW
    EXECUTE FUNCTION atualizar_timestamp();

-- Calcular subtotal automaticamente
CREATE TRIGGER tr_itens_calcular_subtotal
    BEFORE INSERT OR UPDATE ON itens_venda
    FOR EACH ROW
    EXECUTE FUNCTION calcular_subtotal_item();

-- Atualizar total da venda automaticamente
CREATE TRIGGER tr_atualizar_total_venda
    AFTER INSERT OR UPDATE OR DELETE ON itens_venda
    FOR EACH ROW
    EXECUTE FUNCTION atualizar_total_venda();

-- Controlar estoque nas vendas
CREATE TRIGGER tr_controlar_estoque_venda
    BEFORE UPDATE ON vendas
    FOR EACH ROW
    EXECUTE FUNCTION controlar_estoque_venda();

-- Histórico de estoque
CREATE TRIGGER tr_historico_estoque_livros
    AFTER UPDATE ON livros
    FOR EACH ROW
    EXECUTE FUNCTION registrar_historico_estoque();

-- Log de operações
CREATE TRIGGER tr_log_clientes
    AFTER INSERT OR UPDATE OR DELETE ON clientes
    FOR EACH ROW
    EXECUTE FUNCTION log_operacao();

CREATE TRIGGER tr_log_vendas
    AFTER INSERT OR UPDATE OR DELETE ON vendas
    FOR EACH ROW
    EXECUTE FUNCTION log_operacao();

CREATE TRIGGER tr_log_livros
    AFTER INSERT OR UPDATE OR DELETE ON livros
    FOR EACH ROW
    EXECUTE FUNCTION log_operacao();

-- ==========================================
-- VIEWS
-- ==========================================

-- View: Vendas completas
CREATE OR REPLACE VIEW vw_vendas_completas AS
SELECT 
    v.id as venda_id,
    v.data_venda,
    v.status,
    v.forma_pagamento,
    v.valor_total,
    v.observacoes,
    c.id as cliente_id,
    c.nome as cliente_nome,
    c.email as cliente_email,
    c.telefone as cliente_telefone,
    c.cpf as cliente_cpf,
    COUNT(iv.id) as total_itens,
    SUM(iv.quantidade) as total_produtos
FROM vendas v
JOIN clientes c ON v.cliente_id = c.id
LEFT JOIN itens_venda iv ON v.id = iv.venda_id
GROUP BY v.id, c.id;

-- View: Itens de venda detalhados
CREATE OR REPLACE VIEW vw_itens_venda_detalhados AS
SELECT 
    iv.id as item_id,
    iv.venda_id,
    iv.quantidade,
    iv.preco_unitario,
    iv.desconto,
    iv.subtotal,
    l.id as livro_id,
    l.titulo as livro_titulo,
    l.autor as livro_autor,
    l.isbn as livro_isbn,
    l.categoria,
    v.data_venda,
    v.status as venda_status,
    c.nome as cliente_nome
FROM itens_venda iv
JOIN livros l ON iv.livro_id = l.id
JOIN vendas v ON iv.venda_id = v.id
JOIN clientes c ON v.cliente_id = c.id;

-- View: Estoque
CREATE OR REPLACE VIEW vw_relatorio_estoque AS
SELECT 
    l.id,
    l.titulo,
    l.autor,
    l.isbn,
    l.categoria,
    l.editora,
    l.ano_publicacao,
    l.quantidade_estoque,
    l.preco,
    (l.quantidade_estoque * l.preco) as valor_estoque,
    CASE 
        WHEN l.quantidade_estoque = 0 THEN 'SEM_ESTOQUE'
        WHEN l.quantidade_estoque <= 5 THEN 'CRITICO'
        WHEN l.quantidade_estoque <= 15 THEN 'BAIXO'
        WHEN l.quantidade_estoque <= 50 THEN 'NORMAL'
        ELSE 'ALTO'
    END as nivel_estoque,
    l.ativo,
    l.data_ultima_atualizacao
FROM livros l
WHERE l.ativo = true;

-- View: Top clientes
CREATE OR REPLACE VIEW vw_top_clientes AS
SELECT 
    c.id,
    c.nome,
    c.email,
    c.cpf,
    c.datacadastro,
    COUNT(v.id) as total_compras,
    COUNT(CASE WHEN v.status = 'CONFIRMADA' THEN 1 END) as compras_confirmadas,
    COALESCE(SUM(CASE WHEN v.status = 'CONFIRMADA' THEN v.valor_total END), 0) as valor_total_gasto,
    COALESCE(AVG(CASE WHEN v.status = 'CONFIRMADA' THEN v.valor_total END), 0) as ticket_medio,
    MAX(CASE WHEN v.status = 'CONFIRMADA' THEN v.data_venda END) as ultima_compra
FROM clientes c
LEFT JOIN vendas v ON c.id = v.cliente_id
WHERE c.ativo = true
GROUP BY c.id
ORDER BY valor_total_gasto DESC;

-- View: Produtos mais vendidos
CREATE OR REPLACE VIEW vw_produtos_mais_vendidos AS
SELECT 
    l.id,
    l.titulo,
    l.autor,
    l.categoria,
    l.preco,
    COUNT(iv.id) as numero_vendas,
    SUM(iv.quantidade) as total_vendido,
    SUM(iv.subtotal) as receita_total,
    MAX(v.data_venda) as ultima_venda
FROM livros l
JOIN itens_venda iv ON l.id = iv.livro_id
JOIN vendas v ON iv.venda_id = v.id
WHERE v.status = 'CONFIRMADA' AND l.ativo = true
GROUP BY l.id
ORDER BY total_vendido DESC;

-- View: Dashboard
CREATE OR REPLACE VIEW vw_dashboard AS
SELECT 
    (SELECT COUNT(*) FROM vendas WHERE status = 'CONFIRMADA' AND data_venda::date = CURRENT_DATE) as vendas_hoje,
    (SELECT COUNT(*) FROM vendas WHERE status = 'CONFIRMADA' AND EXTRACT(MONTH FROM data_venda) = EXTRACT(MONTH FROM CURRENT_DATE)) as vendas_mes,
    (SELECT COALESCE(SUM(valor_total), 0) FROM vendas WHERE status = 'CONFIRMADA' AND data_venda::date = CURRENT_DATE) as faturamento_hoje,
    (SELECT COALESCE(SUM(valor_total), 0) FROM vendas WHERE status = 'CONFIRMADA' AND EXTRACT(MONTH FROM data_venda) = EXTRACT(MONTH FROM CURRENT_DATE)) as faturamento_mes,
    (SELECT COUNT(*) FROM clientes WHERE ativo = true) as clientes_ativos,
    (SELECT COUNT(*) FROM livros WHERE ativo = true) as livros_ativos,
    (SELECT COUNT(*) FROM livros WHERE quantidade_estoque <= 5 AND ativo = true) as livros_estoque_baixo,
    (SELECT COUNT(*) FROM vendas WHERE status = 'PENDENTE') as vendas_pendentes;

-- ==========================================
-- DADOS INICIAIS
-- ==========================================

-- Categorias
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
('Literatura Infantil'),
('Negócios'),
('Culinária'),
('Arte'),
('Religião'),
('Saúde')
ON CONFLICT (nome) DO NOTHING;

-- Clientes
INSERT INTO clientes (nome, email, telefone, cpf, endereco, datanascimento) VALUES
('João Silva Santos', 'joao.silva@email.com', '(83) 99999-1111', '123.456.789-01', 'Rua das Flores, 123 - Centro, João Pessoa/PB', '1990-05-15'),
('Maria Santos Oliveira', 'maria.santos@email.com', '(83) 99999-2222', '987.654.321-09', 'Av. Principal, 456 - Bairro Novo, João Pessoa/PB', '1985-08-22'),
('Pedro Oliveira Costa', 'pedro.oliveira@email.com', '(83) 99999-3333', '111.222.333-44', 'Rua da Paz, 789 - Vila Esperança, João Pessoa/PB', '1992-12-03'),
('Ana Paula Ferreira', 'ana.paula@email.com', '(83) 99999-4444', '555.666.777-88', 'Rua do Sol, 321 - Tambaú, João Pessoa/PB', '1988-03-10')
ON CONFLICT (cpf) DO NOTHING;

-- Livros
INSERT INTO livros (titulo, autor, isbn, preco, quantidade_estoque, categoria, editora, ano_publicacao) VALUES
('Dom Casmurro', 'Machado de Assis', '978-85-359-0277-5', 29.90, 50, 'Literatura Brasileira', 'Companhia das Letras', 1899),
('Clean Code', 'Robert C. Martin', '978-0-13-235088-4', 89.90, 25, 'Tecnologia', 'Prentice Hall', 2008),
('O Pequeno Príncipe', 'Antoine de Saint-Exupéry', '978-85-359-0624-7', 24.90, 100, 'Literatura Infantil', 'Agir', 1943),
('1984', 'George Orwell', '978-85-359-0336-9', 34.90, 75, 'Ficção Científica', 'Companhia das Letras', 1949),
('O Cortiço', 'Aluísio Azevedo', '978-85-359-0445-8', 27.50, 40, 'Literatura Brasileira', 'Ática', 1890)
ON CONFLICT (isbn) DO NOTHING;

-- Vendas de exemplo
INSERT INTO vendas (cliente_id, forma_pagamento, status, observacoes) VALUES
(1, 'PIX', 'CONFIRMADA', 'Pagamento à vista'),
(2, 'CARTAO_CREDITO', 'CONFIRMADA', 'Parcelado em 3x'),
(3, 'DINHEIRO', 'PENDENTE', 'Aguardando pagamento')
ON CONFLICT DO NOTHING;

-- Itens de venda
INSERT INTO itens_venda (venda_id, livro_id, quantidade, preco_unitario, desconto) VALUES
(1, 1, 2, 29.90, 5.00),
(1, 3, 1, 24.90, 0.00),
(2, 2, 1, 89.90, 10.00)
ON CONFLICT DO NOTHING;

-- ==========================================
-- COMENTÁRIOS
-- ==========================================

COMMENT ON TABLE categorias IS 'Categorias de livros';
COMMENT ON TABLE clientes IS 'Cadastro de clientes';
COMMENT ON TABLE livros IS 'Catálogo de livros';
COMMENT ON TABLE vendas IS 'Registro de vendas';
COMMENT ON TABLE itens_venda IS 'Itens de cada venda';

-- ==========================================
-- FINALIZAÇÃO
-- ==========================================

ANALYZE;

DO $$
BEGIN
    RAISE NOTICE '==========================================';
    RAISE NOTICE 'SISTEMA DE LIVRARIA - COMPATÍVEL COM GO';
    RAISE NOTICE '==========================================';
    RAISE NOTICE 'Tabelas: %', (SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE');
    RAISE NOTICE 'Views: %', (SELECT count(*) FROM information_schema.views WHERE table_schema = 'public');
    RAISE NOTICE 'Triggers: %', (SELECT count(*) FROM information_schema.triggers WHERE trigger_schema = 'public');
    RAISE NOTICE '==========================================';
    RAISE NOTICE 'Sistema pronto! 🚀';
    RAISE NOTICE '==========================================';
END $$;