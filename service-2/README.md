#service-2

Таблицы:

task: {
    id : bigserial,
    name : varchar,
    description varchar,
    price int

}

history: {
    account_id  int,
    user_id int,
    task_id int
}

API:
GET: /history/user_id ->  список истории, решеных задач по uid
GET: /task -> список всех задач
PUT: /task/{id}/price -> обновление цены задачи по  taskID
POST: /task/{id}/solve -> решение задачи, обновление долга в таблице - счет
