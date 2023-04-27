## Описание

Домашняя работа в рамках курса SOA 2023 ФКН ВШЭ

Представляет из себя приложение на go для тестирование разных форматов на скорость сериализации/десериализации и размера данных в сериализованном виде.

Тестируемые форматы/использумые библиотеки:
* Native gob - https://pkg.go.dev/encoding/gob
* XML - https://pkg.go.dev/encoding/xml
* JSON - https://pkg.go.dev/encoding/json
* Google Protocol Buffers - https://pkg.go.dev/google.golang.org/protobuf@v1.30.0
* Apache Avro - https://pkg.go.dev/github.com/hamba/avro/v2@v2.8.1
* YAML - https://pkg.go.dev/gopkg.in/yaml.v2@v2.4.0#section-readme
* Message Pack - https://pkg.go.dev/github.com/vmihailenco/msgpack/v5@v5.3.5

Для каждого формата поднимается свой контейнер + поднимается прокси для единой точки доступа. Общение происходит по UDP.
## Сборка и запуск

```
cd build
docker compose up
```

<details>
    <summary> Локальная сборка </summary>
    Для сборки в обход DockerHub нужно раскомментировать секции build в docker-compose.yml и выполнить
    docker compose build
</details>

## Формат запросов

В контейнеры для каждого формата можно сходить с запросом ```get_result``` по UDP (2001-2007 порты):

```
$ echo -n "get_result" | nc -u -w 1 127.0.0.1 2004
GoogleProtocolBuffers-1277-14.748µs-13.386µs
```

Есть прокси на 2000 порту, через который можно получить ответ с запросом ```get_result {format}```:
```
$ echo -n "get_result avro" | nc -u -w 1 127.0.0.1 2000
ApacheAvro-1038-10.137µs-6.265µs
```

Форматы в запросе: ```native, xml, json, protobuf, avro, yaml, message_pack```

Также через прокси можно получить все результаты запросом ```get_result all```:
```
$ echo -n "get_result all" | nc -u -w 2 127.0.0.1 2000
ApacheAvro-1038-5.99µs-8.981µs
Native(gob)-2018-12.92µs-11.54µs
GoogleProtocolBuffers-1277-6.896µs-23.613µs
JSON-3556-21.228µs-58.782µs
MessagePack-2690-35.317µs-44.862µs
XML-7907-103.782µs-378.674µs
YAML-4327-269.368µs-346.812µs
```

## Подробности для проверки

Реализован продвинутый вариант с прокси + публикация на DockerHub:

https://hub.docker.com/repository/docker/azuremint/serializers-proxy/general

https://hub.docker.com/repository/docker/azuremint/serialize-benchmarker/general

Очень продвинутый вариант не реализован. Хоть и есть ```get_result all```, но он выполнятся не multicast запросом, а через обычные параллельные запросы в контейнеры.
