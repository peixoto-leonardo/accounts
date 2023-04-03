GRANT ALL PRIVILEGES ON DATABASE bank TO dev;

CREATE TABLE accounts (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    cpf CHAR(11) UNIQUE NOT NULL,
    balance BIGINT NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL
);

CREATE TYPE transaction_type_enum AS ENUM('WITHDRAW', 'DEPOSIT');

CREATE TABLE transactions (
    id UUID PRIMARY KEY NOT NULL,
    account_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    type transaction_type_enum,
    FOREIGN KEY (account_id)
        REFERENCES accounts (id)
);