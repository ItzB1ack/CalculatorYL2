package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/calculate", calculateHandler).Methods("POST")
	r.HandleFunc("/api/v1/expressions", expressionsHandler).Methods("GET")
	r.HandleFunc("/api/v1/expressions/{id}", expressionHandler).Methods("GET")
	r.HandleFunc("/internal/task", taskHandler).Methods("GET")
	r.HandleFunc("/internal/task", taskResultHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Выражение принято для вычисления")
}

func expressionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Список выражений")
}

func expressionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Выражение по идентификатору")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Задача для выполнения")
}

func taskResultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Результат обработки данных")
}
