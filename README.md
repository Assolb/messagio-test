# Микросервис обработки сообщений

## Описание

Этот проект представляет собой микросервис, разработанный на языке Go, предназначенный для обработки сообщений через HTTP API. Он сохраняет сообщения в базе данных PostgreSQL и отправляет их в очередь Kafka для дальнейшей обработки. Также предоставляется API для получения статистики по обработанным сообщениям.

## Запуск проекта

### Предварительные требования

Убедитесь, что у вас установлены [Docker](https://www.docker.com/get-started) и [Docker Compose](https://docs.docker.com/compose/install/).

### Шаги по запуску

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/Assolb/messagio-test.git
   cd messagio-test

2. **Сборка и запуск проекта**

   Убедитесь, что у вас установлен Docker и Docker Compose. Затем выполните команду для запуска контейнеров.

   ```bash
   docker-compose up --build
   
## Доступ к API

Проект развернут и доступен для тестирования на сервере по следующему адресу:

[http://dhazov19.fvds.ru/](http://dhazov19.fvds.ru/)

API включает следующие эндпоинты:
- `POST /api/v1/message/add` — Отправка сообщений.
- `GET /api/v1/message/stats` — Получение статистики по обработанным сообщениям.
