
CREATE TABLE IF NOT EXISTS  task
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(255),
    price INTEGER coalesce(0),
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