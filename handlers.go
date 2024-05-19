package main

import (
	"encoding/json"
	"net/http"
)

// UsersHandler для отдачи списка пользователей
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	json.NewEncoder(w).Encode(users)
}

// ProductsHandler для отдачи списка продуктов
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 1200.00},
		{ID: 2, Name: "Smartphone", Price: 800.00},
	}
	json.NewEncoder(w).Encode(products)
}

// HealthCheckHandler для проверки состояния сервиса
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
