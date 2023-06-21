how to run:
1.docker-compose up
2. ./run-sql.sh

service1 url: localhost:9090/v1/
service2 url: localhost:6969/v1/
psql: localhost:5432





GET: /history/user_id ->  список истории, решеных задач по uid
GET: /task -> список всех задач
PUT: /task/{id}/price -> обновление цены задачи по  taskID
POST: /task/{id}/solve -> решение задачи, обновление долга в таблице - счет

POST: /user -> Создание нового юзера
GET: /account/:userId -> Счет юзера по userID
PUT: /account/debt/decrease - Уменьшение долга юзера
PUT: /account/debt/increase - Увелечние долга юзера
