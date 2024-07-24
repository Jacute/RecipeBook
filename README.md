# RecipeBook

RecipeBook — это веб-приложение для создания, хранения и просмотра рецептов. Проект включает в себя API, веб-приложение и базу данных, работающую на MySQL.

## Оглавление

1. [Требования](#требования)
2. [Установка](#установка)
3. [Запуск проекта](#запуск-проекта)
4. [Использование](#использование)
5. [Структура проекта](#структура-проекта)
6. [API](#api)
7. [Авторизация и аутентификация](#авторизация-и-аутентификация)
8. [База данных](#база-данных)

## Требования

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Установка

1. Склонируйте репозиторий:
    ```bash
   git clone https://github.com/Jacute/RecipeBook.git
   cd RecipeBook
   ```

2. Внесите данные БД в .env_db и создайте jwt ключи с помощью gen_jwt_keys.sh:

   ```bash
   nano .env_db # change this
   chmod +x gen_jwt_keys.sh
   ./gen_jwt_keys.sh
   ```

## Запуск проекта

1. Запустите контейнеры с помощью Docker Compose:
   ```bash
   docker-compose up --build
   ```

2. Откройте браузер и перейдите по адресу:
   ```http
   http://127.0.0.1
   ```

## Использование

- Перейдите на главную страницу, чтобы просмотреть список публичных рецептов.
- Зарегистрируйтесь или войдите в систему, чтобы создавать новые рецепты.

## Структура проекта

Проект состоит из нескольких основных компонентов:

- `api/`: Реализация API для управления рецептами.
- `webapp/`: Веб-приложение для взаимодействия с пользователем.
- `jwt_keys/`: JWT ключи для авторизации и взаимодействия с API.
- `nginx/`: Конфигурационные файлы для NGINX.

## API

### Эндпоинты

- `GET /api/recipes`: Получить список всех публичных рецептов.
- `GET /api/recipes/:id`: Получить рецепт по ID.
- `GET /api/images/:filename`: Получить изображение рецепта.

### Авторизация и аутентификация

- Регистрация:
  ```http
  POST /register
  Body: {
      "username": "string",
      "email": "string",
      "password": "string"
  }
  ```

- Вход:
  ```http
  POST /login
  Body: {
      "username": "string",
      "password": "string"
  }
  ```

- Выход:
  ```http
  GET /logout
  ```
### Тесты API

```bash
cd api
sudo ./run_tests.sh
```

## База данных

### Таблицы

- `users`: Таблица пользователей.
- `recipes`: Таблица рецептов.

### Миграции

При инициализации базы данных автоматически создаются необходимые таблицы и заполняются тестовыми данными из `init.sql`.

### Пример конфигурации базы данных (`.env_db`)

```ini
DATABASE_HOSTNAME=recipebook_db
MYSQL_ROOT_PASSWORD=root_password
MYSQL_DATABASE=recipebook
MYSQL_PASSWORD=user_password
MYSQL_USER=user

```

