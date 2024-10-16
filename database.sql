-- Tabela de usuários (users)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    plan_id INT, -- Campo para associar o plano ao usuário
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100),
    age NUMERIC,
    status BOOLEAN NOT NULL,
    CONSTRAINT fk_plan FOREIGN KEY (plan_id) REFERENCES plans(id) -- Chave estrangeira para a tabela de planos
);

-- Tabela de planos (plans)
CREATE TABLE plans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- Tabela de pagamentos (payments)
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    month VARCHAR(20) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT FALSE, -- Status do pagamento (pago ou não)
    payment_date TIMESTAMP, -- Data do pagamento
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) -- Chave estrangeira para a tabela de usuários
);
