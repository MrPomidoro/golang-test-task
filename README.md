# SoftWeather. Тестовое задание в команду финтеха Golang.

## Проблема

В ходе работы зачастую приходится искать нестандартные и оптимизированные пути решения задач. Этим тестовым заданием мы хотим оценить твое умение выполнять задачу, имея неполные данные о ней, а также умение самостоятельно принимать решения и соблюдать качество кода, поэтому пиши код так, как бы ты его писал, работая над энтерпрайз проектом.

Представим, что у ты учишься в университете на программиста, и к тебе постоянно обращаются люди с учебного потока, которые ничего не понимают в алгоритмах, с просьбой помочь им с решением за деньги. В один прекрасный день тебе пришла в голову идея, что надо создать приложение, которое делало бы за тебя всю работу, а тебе оставалось только собирать деньги с обратившихся к тебе людей.

## Задача

Реализовать сервис, который предоставляет следующий функционал:

1. Создание нового студента, используя его ФИО, номер группы, а также email, например, Штирлиц Иван Васильевич ИП-394 vasya@mail.ru. Для удобства дальнейших запросов, каждому пользователю можно присваивать логин, который будет в роли ярлыка для этого студента.
2. Изменение долга студента (уменьшение и увеличение). Максимальный долг 1000 рублей. При достижении максимального долга, приложение не должно давать ответы на новые запросы решения алгоритмических задач.
3. Получение ответа на алгоритмическую задачу. У каждой задачи есть своя стоимость для получения ответа, поэтому необходимо изменять долг студента, для которого надо получить ответ (не забудь про максимальный долг).

**Единственная алгоритмическая задача, на которую можно получить ответ и надо реализовать:**

```
Дан неотсортированный массив из N чисел от 1 до N,
при этом несколько чисел из диапазона [1, N] пропущено, 
а некоторые присутствуют дважды.

Найти все пропущенные числа.
```

## Пути разработки сервиса

Сервис можно реализовать одним из следующих путей:

1. REST API.
2. CLI-утилита.

Выбранный путь необходимо отметить в README файле.

## Требования к сервису

1. Язык Golang. Можно использовать любые фреймворки и библиотеки.
2. В качестве БД использовать PostgreSQL.
3. Весь код должен быть выложен на Github или Gitlab с README файлом, в котором содержится инструкция по запуску и примерами запросов/ответов.

При возникновении вопросов по ТЗ оставляем принятие решения за кандидатом (в таком случае в README файле к проекту должен быть указан список вопросов с которыми кандидат столкнулся и каким образом он их решил).

Разработка интерфейса в браузере НЕ ТРЕБУЕТСЯ. Для тестирования можно использовать любой удобный инструмент. Например: в терминале через curl или Postman.

Систему авторизации и аутентификации реализовывать НЕ ТРЕБУЕТСЯ. Предполагается, что проект будет использоваться на личном компьютере.

## Будет плюсом

Если ты дополнительно сделаешь что-то из этого списка, то ты заметно выделишься среди других кандидатов!

1. Покрытие кода тестами.
2. Отправка email на почту студента, в котором будет содержаться ответ на задачу или уведомление о том, что у этого студента превышен долг.
3. Использование docker и docker-compose для поднятия и развертывания рабочей среды.
4. Примитивное CI/CD (достаточно только сборки и авто-тестов).
5. Makefile.

# Исполнение

## Для запуска

## Rest API документация OpenAPI Swagger
```bash
  http://localhost:8080/swagger/
```

## Создание пользователя

### 1. Необходимо отправить POST запрос по url например с помощью CURL:
```bash
curl -X 'POST' \
  'http://localhost:8080/api/v1/students/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "qweqw@qwe.com",
  "full_name": "asdsad Dmitrij asdasd",
  "group_num": "ANTM-23m",
  "username": "asdqw"
}'
```
### Тело ответа при успешном создании

```json
{
  "status": "success",
  "message": "Success create student",
  "data": {
    "id": "5218a4ac-f8f2-47bd-b2dd-11b3041766c8",
    "full_name": "asdsad Dmitrij asdasd",
    "group_num": "ANTM-23m",
    "email": "qweqw@qwe.com",
    "username": "asdqw",
    "verify_email": false,
    "create_at": "2023-07-07T13:46:26.912952503+03:00",
    "update_at": null
  }
}
```

### При не удачном (на примере создан cтудент с существующим username)

```json
{
  "status": "Internal Server Error",
  "error": "Repository.Create: failed to query: ERROR: duplicate key value violates unique constraint \"students_username_key\" (SQLSTATE 23505)"
}
```

## Для получения ответа на задачу, сначала надо выбрать задвчу из общегго списка

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v1/tasks/' \
  -H 'accept: application/json'
 ```

```json

{
  "status": "success",
  "message": "Success get all tasks",
  "data": [
    {
      "id": "48acfe1a-9bcb-41fd-b5d6-e34f14c10f87",
      "description": "Дан неотсортированный массив из N чисел от 1 до N,\nпри этом несколько чисел из диапазона [1, N] пропущено,\nа некоторые присутствуют дважды.\n\nНайти все пропущенные числа.",
      "cost": 100,
      "create_at": "2023-07-07T14:06:48.940309+03:00",
      "update_at": null
    }
  ]
}
```
### Если интересует конкретная задача, то можно найти по ее UUID

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v1/tasks/48acfe1a-9bcb-41fd-b5d6-e34f14c10f87' \
  -H 'accept: application/json'
```

```json

{
  "status": "success",
  "message": "Success get task",
  "data": {
    "id": "48acfe1a-9bcb-41fd-b5d6-e34f14c10f87",
    "description": "Дан неотсортированный массив из N чисел от 1 до N,\nпри этом несколько чисел из диапазона [1, N] пропущено,\nа некоторые присутствуют дважды.\n\nНайти все пропущенные числа.",
    "cost": 100,
    "create_at": "2023-07-07T14:06:48.940309+03:00",
    "update_at": null
  }
}
```

### На основании полученной задачи мы имеем цены за задачи, которые мы отправляем в метод списания денег и получения ответа на email

```bash
curl -X 'POST' \
  'http://localhost:8080/api/v1/jobs/slow_task_missing_numbers' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "amount": 100,
  "nums": [4, 3, 2, 7, 8, 2, 3, 1],
  "username": "SDA"
}'
```

### Так же мы можем пополнить счет определенного пользователя
```bash
curl -X 'POST' \
  'http://localhost:8080/api/v1/jobs/add_credit?amount=12&username=SDA' \
  -H 'accept: application/json' \
  -d ''
```

# ВАЖНО

#### 1. Перед запуском приложения, определить все environment переменныe
#### 2. Найти переменные можно в Dokerfile
#### 3. Для упрощения пользования, был описан Makefile
#### 4. Так же описан конфигурационный файл для air

# Вопросы, которые не вошли в реализацию

#### 1. Необходимо расписать полностью все методы, для переиспользования их на адмиской стороне;
#### 2. Сделать распределение ролей (админ, обычный пользователь);
#### 3. Сделать верификацию email, как самое просто решение использовать redis для кеширования;
#### 4. Авторизация по логину, при верификации email;
#### 5. Добавить варианты решение кодом, а не просто ответом (генерация для 2-3х языков)

