package main

import (
	"Backend/internal"
	"Backend/internal/postgres"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"

	"net/http"
	"os"
)

type application struct {
	Store internal.DataStorage
}

const defaultPort = "8080"

func main() {

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = defaultPort
	}

	store, db := internal.NewDataStorage()
	fmt.Println("Успешный запуск базы данных postgres")
	defer postgres.CloseDB(db)

	app := &application{
		Store: store,
	}

	router := mux.NewRouter()
	router.HandleFunc("/status", app.getStatus).Methods("GET")
	router.HandleFunc("/status", app.addStatus).Methods("POST")

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},  // Разрешить запросы с фронтенда
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // Разрешить методы
		AllowedHeaders:   []string{"*"},                      // Разрешить все заголовки
		AllowCredentials: true,                               // Разрешить передачу кук и авторизационных данных
	})

	// Оберните роутер в CORS
	handler := c.Handler(router)

	log.Println("backend-сервис запущен на :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
