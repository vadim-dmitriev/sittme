# Приложение для управления состоянием объектов "Трансляция"

## Описание

Приложение представляет из себя веб-сервер, работающий по протоколу HTTP. API сервера спроекторован в соответствии с принципами [REST]( https://en.wikipedia.org/wiki/Representational_state_transfer) и [JSON API](http://jsonapi.org/).

Сервис предоставляет возможности:
- создание объекта трансляции
- удаление объекта трансляции
- получение списка существующих трансляций
- изменение состояния трансляции

## Запуск сервиса

### Требования

Для запуска необходима версия [Golang](https://golang.org/) >= 1.11
```bash
$ go version
go version go1.11.5 linux/amd64
```

### Установка

```bash
go get github.com/vadim-dmitriev/sittme
```

### Запуск

```bash
cd $GOPATH/src/github.com/vadim-dmitriev/sittme
$ go build . && ./sittme
2020/06/21 18:52:11 Service started on localhost:8080
```

## Параметры запуска приложения

Приложение имеет два основных источника параметров запуска:
- Переменные окружения 
  - PORT содержит целочисленное значение порта, на котором должно работать приложение
  - RTIMEOUT и WTIMEOUT содержат целочисленые значения времени (в секундах) читения и записи запроса/ответа соответственно
  - TIMEOUT содержит целочисленное значение таймера (в секундах) перехода из состояния Interrupted в Finished

- Аргументы запуска
Аргументы запуска можно посмотреть, выполнив `./sittme --help`. Они содержат такой же набор настраиваемых параметров, как и переменне окрежения.

> NOTE: Значения переменных окружения являются приоритетными.

## API

### `POST /api/v1/streams`

Создание трансляции. Трансляции назначется уникальный идентификатор.

```bash
$ http POST :8080/api/v1/streams
HTTP/1.1 201 Created
Content-Length: 155
Content-Type: application/json
Date: Sun, 21 Jun 2020 16:19:27 GMT
Server: fasthttp

{
    "data": {
        "attributes": {
            "date_modified": "2020-06-21T19:19:27.365443063+03:00",
            "state": "created"
        },
        "id": "bd1fae7a-8b54-418f-85b1-461ed9c1d781"
    },
    "error": null
}

```

### `GET /api/v1/sreams`

Получение списка всех существующих трансляций.

```bash
$ http GET :8080/api/v1/streams
HTTP/1.1 200 OK
Content-Length: 290
Content-Type: application/json
Date: Sun, 21 Jun 2020 16:20:44 GMT
Server: fasthttp

{
    "data": [
        {
            "attributes": {
                "date_modified": "2020-06-21T19:19:27.365443063+03:00",
                "state": "created"
            },
            "id": "bd1fae7a-8b54-418f-85b1-461ed9c1d781"
        },
        {
            "attributes": {
                "date_modified": "2020-06-21T19:20:42.920917255+03:00",
                "state": "created"
            },
            "id": "1955bf05-28b4-4dc4-8a2b-40b329ac4a34"
        }
    ],
    "error": null
}
```

### `DELETE /api/v1/streams/{uuid}`

Удаление трансляции с заданным уникальным идентификатором `uuid`.

```bash
$ http DELETE :8080/api/v1/streams/ad7b989e-f36e-4570-967c-49ce807386c8
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Sun, 21 Jun 2020 16:25:20 GMT
Server: fasthttp


```

### `PATCH /api/v1/streams/{uuid}`

Изменение состояния трансляции с заданным уникальным идентификатором `uuid`. В теле запроса необходимо передать JSON следующего формата:
```json
{
    "state": "{new_state}"
}
```

```bash
$ echo '{"state": "active"}' | http PATCH :8080/api/v1/streams/d2a444de-d49d-4b73-aa4c-fba5bd34a8ea
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Sun, 21 Jun 2020 16:28:55 GMT
Server: fasthttp


```