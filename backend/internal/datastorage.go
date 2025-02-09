package internal

import (
	"Backend/internal/model"
	"Backend/internal/postgres"
	"database/sql"
	"fmt"
	"os"
	"time"
)

type DataStorage interface {
	AddStatus(ip string, alive bool, checked, lastSuccess time.Time) error
	GetAllStatuses() ([]*model.ContainerStatus, error)
}

const (
	defaultDbPort     = "5432"
	defaultDbHost     = "localhost"
	defaultDbUser     = "postgres"
	defaultDbPassword = ""
	defaultDbName     = "postgres"
)

func NewDataStorage() (DataStorage, *sql.DB) {
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		port = defaultDbPort
	}
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = defaultDbHost
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = defaultDbUser
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = defaultDbPassword
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		dbname = defaultDbName
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	store, err := postgres.OpenDB(connStr)
	if err != nil {
		fmt.Println(err)
	}
	return store, store.DB
}
