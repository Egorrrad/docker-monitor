# Мониторинг Docker-контейнеров

Система для отслеживания состояния Docker-контейнеров через REST API и веб-интерфейс.

## 🚀 Функциональность

- Получение IP-адресов Docker-контейнеров
- Пинг контейнеров с заданным интервалом
- Сохранение данных в PostgreSQL
- Веб-интерфейс с таблицей статусов

## 🛠 Технологии

**Backend:**  
![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)  
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-4169E1?logo=postgresql)  
![Gorilla/Mux](https://img.shields.io/badge/Gorilla%2FMux-1.8+-7F479B)

**Frontend:**  
![React](https://img.shields.io/badge/React-18+-61DAFB?logo=react)  
![Ant Design](https://img.shields.io/badge/Ant%20Design-5+-0170FE)  
![Axios](https://img.shields.io/badge/Axios-1.3+-5A29E4)

**Инфраструктура:**  
![Docker](https://img.shields.io/badge/Docker-20.10+-2496ED?logo=docker)  
![Docker Compose](https://img.shields.io/badge/Compose-2.12+-2496ED)

## ⚙️ Установка и запуск

1. Клонировать репозиторий:
```bash
git clone https://github.com/Egorrrad/docker-monitor.git
cd docker-docker-monitor
```

2. Запустить сервисы:
```bash
docker-compose up --build
```

3. Открыть в браузере:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8200/status

4. Остановить сервисы после использования:
```bash
docker-compose up --build
```

## 📂 Структура проекта

```
docker-monitor/
├── backend/               # Go-сервис
│   ├── Dockerfile
│   ├── handlers.go
│   ├── main.go
│   └── internal/
├── frontend/              # React-интерфейс
│   ├── Dockerfile
│   ├── public/
│   ├── src/
│   └── package.json
├── pinger/                # Сервис мониторинга
│   ├── Dockerfile
│   └── main.go
├── docker-compose.yml
└── README.md
```
