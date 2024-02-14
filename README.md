# Backlog REST API service

Пет-проект на один вечер - CRUD бэклога с REST API и SQLite.

## Запуск

Требуется установить переменную окружения `CONFIG_PATH` (к примеру на `./config/local.yml`).

```
$ go run cmd/backlog/main.go
```

## Использование

Задача имеет вид
```
{
    "id": int
    "text": string
    "status": "To do" | "In progress" | "Done" 
}
```

Ручки:
- `GET /` возвращает все задачи
- `POST /` добавляет новую задачу со статусом "To do", 
требует поле `text` в теле запроса (к примеру, `{"text": "Finish project"}`)
- `PUT /` переносит задачу в другой статус, требует поля `id` и `status`

## TODO

- Тесты
- Выделить сервисный слой? 
- Подключить swagger