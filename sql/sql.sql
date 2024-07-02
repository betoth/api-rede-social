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
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO golang;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO golang;


CREATE TABLE followers (
    users_id INT NOT NULL, 
    follower_id INT NOT NULL, 

    CONSTRAINT fk_users
        FOREIGN KEY (users_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_followers
        FOREIGN KEY (follower_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    PRIMARY KEY (users_id, follower_id)
);

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE followers TO golang;
GRANT USAGE, SELECT ON SEQUENCE followers_users_id_seq TO golang;
GRANT USAGE, SELECT ON SEQUENCE followers_follower_id_seq TO golang;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO golang;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO golang;
GRANT ALL PRIVILEGES ON SCHEMA public TO golang;


CREATE TABLE publications (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id BIGINT NOT NULL,
    likes BIGINT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_author
        FOREIGN KEY (author_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);

GRANT USAGE, SELECT ON SEQUENCE publications_id_seq TO golang;