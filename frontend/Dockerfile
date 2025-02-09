# Используем базовый образ Node.js
FROM node:16-alpine as build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем исходный код
COPY . .

# Собираем приложение в режиме продакшена
RUN npm run build

# Устанавливаем serve для обслуживания статических файлов
RUN npm install -g serve

# Открываем порт 3000
EXPOSE 3000

# Запускаем serve для обслуживания статических файлов
CMD ["serve", "-s", "build", "-l", "3000"]