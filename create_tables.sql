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


insert into users(name, login) values('test','login');
insert into task(name, price) values('task1', 241);
insert into account(user_id) values(1);