# rest-api-go

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 201, 4xx, 500, Header Location: url
PUT /users/:id -- fully update user -- 204/200, 404, 400, 500
PATCH /users/:id -- partially update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400

project struct:

rest-api-go:

    cmd -- точки входа в приложения, можно создавать подкаталоги с названиями приложений
    internal -- вся внутренняя кухня сервиса (специфичный код - именно для этого приложения, бизнес логика )
    pkg -- код, который мы можем переиспользовать (утилиты, клиенты и пр.)