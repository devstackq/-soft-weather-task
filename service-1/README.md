service-1:

Таблицы: 

users: {
    id :  bigserial,
    name : varchar,
    login: varchar
}

account: {
    id : bigserial,
    user_id int,
    debt int,
    balance int
}

API:
POST: /user -> Создание нового юзера
GET: /account/:userId -> Счет юзера по userID
PUT: /account/debt/decrease - Уменьшение долга юзера
PUT: /account/debt/increase - Увелечние долга юзера

how to run:
docker-compose up;
service1 url: localhost:8080/v1/;
service2 url: localhost:9090/v1/;
