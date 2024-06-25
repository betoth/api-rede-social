-- Conectar ao banco de dados padrão 'postgres' para criar o novo banco de dados
\c postgres

-- Criar o banco de dados 'social_network'
CREATE DATABASE social_network;

-- Conectar ao banco de dados 'social_network'
\c social_network

-- Apagar a tabela 'users' se ela já existir
DROP TABLE IF EXISTS users;

-- Criar a tabela 'users'
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick_name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
