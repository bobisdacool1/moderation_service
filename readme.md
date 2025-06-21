# 🛡️ ModerationService

Мне стало скучно жить, и я решил написать сервис по эмуляции модерации.

Kafka вместо базы, память вместо диска, архитектура — **чистая, как яйца кота**.  

---

## 🎯 Цели проекта

- ✅ Управление заявками через REST API
- ✅ Kafka — **единственное хранилище** (никаких БД, только хардкор)
- ✅ In-memory TTL-кэш для "живых" заявок
- ✅ Отделение слоёв: handler → usecase → service → adapter
- ✅ DI через `uber-go/fx`, потому что мы не в пещере

---

## 🧱 Архитектура

```

HTTP (Fiber)
│
Handler
│
Usecase
│
Service
│
├── KafkaAdapter
└── InMemAdapter

````

Каждый слой знает только о нижележащем. Зависимости не текут вверх

---

## 🧪 Жизненный цикл заявки

1. Кто-то стучит в `/moderation` — заявка уходит в Kafka
2. Модератор берёт `/moderation/next` — заявка паркуется в памяти
3. `/approve` или `/decline` — она уходит в нужный топик, а память прощается
4. Если не ответил — TTL её сожрёт

---

## ⚙️ Быстрый старт

```bash
docker-compose up -d    # поднимаем Kafka
go run ./cmd/main.go    # запускаем красоту
````

Порт: `:3000`
Путь к блаженству: `/api/moderation`

---

## 📡 API

| Метод  | Путь                          | Описание                  |
| ------ | ----------------------------- |---------------------------|
| `POST` | `/api/moderation`             | Создать заявку            |
| `GET`  | `/api/moderation/next`        | Взять следующую           |
| `POST` | `/api/moderation/:id/approve` | Подтвердить               |
| `POST` | `/api/moderation/:id/decline` | Отклонить                 |
| `GET`  | `/health`                     | Мы живы? Живее всех живых |

---

## 📦 Конфиг

Файл: `config.yaml`

```yaml
app:
  name: ModerationService
  version: "1.0"

server:
  port: 3000

kafka:
  broker: localhost:9092
  topics:
    - alias: moderation-requests
      topic: moderation-requests
      group_id: moderator-group
```

---

## 🤯 Почему нет БД?

Потому что Kafka — это тоже база.
Заявки лежат в логах, подтверждаются через offset-коммиты.
Если хочешь durability — иди в Postgres. А тут усложняющий себе жизнь **эксперимент**.

---

## 🛠 TODO (на случай, если станет ещё скучнее)

* [ ] Тесты, чтобы не бояться
* [ ] Swagger, чтобы казаться серьёзнее
* [ ] Метрики, чтобы видеть, как оно падает
* [ ] Graceful shutdown Kafka-клиента, потому что Ctrl+C — не всегда милосерден

---

## ✍️ Автор

Aleksey Sinitsyn
*Человек, который зачем-то решил сделать собственный велосипед*

---
