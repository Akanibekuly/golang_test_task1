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
