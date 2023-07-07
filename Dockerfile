# Установка базового образа
FROM golang:latest

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Установка зависимостей
RUN go mod download

# Сборка приложения
RUN go build -o main .

# Определение переменных окружения
ENV LOG_LEVEL="debug"

# Database settings:
ENV POSTGRES_HOST="localhost"
ENV POSTGRES_PORT="5432"
ENV POSTGRES_USER="root"
ENV POSTGRES_PASSWORD="password"
ENV POSTGRES_DB="test"

# Server settings:
ENV SERVER_URL="localhost:8080"

# GIN_MODE settings
ENV GIN_MODE="debug"

# SMTP settings:
ENV SMTP_SENDER="dmitrijsemenkin@gmail.com"
ENV SMTP_HOST="smtp.gmail.com"
ENV SMTP_PORT=587
ENV SMTP_USERNAME="dmitrijsemenkin@gmail.com"
ENV SMTP_PASSWORD="sosiska3"

# Открытие порта
EXPOSE 8080

# Запуск приложения
CMD ["./main"]


