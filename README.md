# Тестовое задание для Golang

## Описание
Написать HTTP сервер для обработки базовых методов (CRUD) таблицы `city`.

## Описание таблицы
* `id` serial (primary key)
* `name` text not null
* `code` text not null
* `country_code` text not null

## СУБД
SQLite или PostgreSQL

## Итоговые HTTP ендпойнты
* GET /cities - список городов
* POST /cities - создать новый город
* GET /cities/{id} - получить один город по id
* PUT /cities/{id} - изменить город по id
* DELETE /cities/{id} - удалить город по id

## Требование к оформлению
* Нужно чтобы хост и порт для http-сервера можно было указывать через переменную среду. Например: HTTP_PORT=127.0.0.1:9090 go run main.go
* Необходимо разместить код проекта в github.com

## Желаемые дополнения
Наличие авто-тестов

## Результат
Ссылка на репозитории в github.com

1. 
 * POST http://127.0.0.1:9090/api/v1/cities
```json
{
    "name": "Almaty",
    "code": "727",
    "country_code": "02"
}
```
response: 
```json
{
    "status": "OK",
    "message": "city with id 4 successfully created"
}
```

2. 
* GET http://127.0.0.1:9090/api/v1/cities
* response:
```json:
{
    "status": "OK",
    "data": [
        {
            "id": 1,
            "name": "Almaty",
            "code": "727",
            "country_code": "02"
        },
        {
            "id": 2,
            "name": "Almaty",
            "code": "727",
            "country_code": "02"
        },
        {
            "id": 3,
            "name": "Almaty",
            "code": "727",
            "country_code": "02"
        },
        {
            "id": 4,
            "name": "Almaty",
            "code": "727",
            "country_code": "02"
        }
    ]
}
```
3. GET http://127.0.0.1:9090/api/v1/cities/3
* response: 
```json:
{
    "status": "OK",
    "data": {
        "id": 3,
        "name": "Almaty",
        "code": "727",
        "country_code": "02"
    }
}
```
5. DELETE http://127.0.0.1:9090/api/v1/cities/4
* response: 
```json
{
    "status": "OK",
    "message": "city with id 4 succesfully deleted"
}
```
6. PUT http://127.0.0.1:9090/api/v1/cities/3
```json
{
    "name": "Astana",
    "code": "7172",
    "country_code": "01"
}
```
* response: 
```json
{
    "status": "OK",
    "message": "city with id 3 succesfully deleted"
}
```