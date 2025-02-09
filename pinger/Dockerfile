# Используем базовый образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY . .

# Устанавливаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o pinger .

# Запускаем приложение
CMD ["./pinger"]