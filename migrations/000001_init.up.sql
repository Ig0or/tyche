CREATE TABLE IF NOT EXISTS accounts (
    id INT NOT NULL PRIMARY KEY,
    account_id UUID NOT NULL,
    email VARCHAR(50) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    password VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT NOT NULL PRIMARY KEY,
    operation VARCHAR(20) NOT NULL,
    type VARCHAR(20) NOT NULL,
    amount DECIMAL NOT NULL,
    to_account INT REFERENCES accounts,
    from_account INT REFERENCES accounts,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
