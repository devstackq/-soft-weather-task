CREATE DATABASE  testdb;

CREATE TABLE IF NOT EXISTS  users
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100),
    login VARCHAR(60)
    );

CREATE TABLE IF NOT EXISTS  account
(
    id BIGSERIAL PRIMARY KEY,
    debt INTEGER coalesce(0),
    balance INTEGER coalesce(0),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
    );