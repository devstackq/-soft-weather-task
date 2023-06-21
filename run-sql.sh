#!/bin/bash

sudo docker exec -i postgresdb psql -U postgres -d postgres << EOF
CREATE TABLE IF NOT EXISTS  users
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100),
    login VARCHAR(60)
    );

CREATE TABLE IF NOT EXISTS  account
(
    id BIGSERIAL PRIMARY KEY,
    debt INTEGER,
    balance INTEGER,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
    );

CREATE TABLE IF NOT EXISTS  task
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(255),
    price INTEGER
    );
CREATE TABLE IF NOT EXISTS  history
(
    id BIGSERIAL PRIMARY KEY,
    account_id INT,
    task_id INT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (task_id) REFERENCES task(id)
    );

insert into task(name, price) values('taska1', 143);
insert into task(name, price) values('taska2', 503);
insert into task(name, price) values('taska3', 87);


EOF