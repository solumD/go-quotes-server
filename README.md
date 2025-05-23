# go-quotes-server
## Использованные библиотеки

Практически во всем проекте старался использовать стандартную библиотеку, но в этих случаях воспользовался сторонними пакетами:
  * Подключение к Postgres - для удобства выбрал pgxpool
  * Получение переменных среды - godotenv
  * Роутинг - gorilla/mux 
  * Генерация моков для тестов - gomock

## Установка проекта и скачивание зависимостей
```bash
 git clone github.com/solumD/go-quotes-server
 cd go-quotes-server/
 make install-deps
 go mod tidy
```

## Запуск (docker-compose обязателен)
Поменять значения в .env-файле, если необходимо.
```dotenv
  PG_DATABASE_NAME=quote
  PG_USER=quote-user
  PG_PASSWORD=quote-password
  PG_PORT=54321
  MIGRATION_DIR=./migrations
  
  PG_DSN="host=localhost port=54321 dbname=quote user=quote-user password=quote-password sslmode=disable"
  MIGRATION_DSN="host=pg port=5432 dbname=quote user=quote-user password=quote-password sslmode=disable"
  
  LOGGER_LEVEL=local # local, dev, prod
  
  SERVER_HOST=localhost
  SERVER_PORT=8080
```

Выполнить в терминале:
```bash
  make build-and-run
```
При вводе команды проект компилируется и запускатеся. БД Postgres поднимается в отдельном docker-контейнере, а также накатываются миграции. 

## Тестирование
Базовые unit-тесты успел написать только для handler-слоя. Для запуска выполнить команду в терминале: 
```bash
  make test-handler
```

