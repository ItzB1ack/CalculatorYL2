package main

import (
	"fmt"
	"github.com/ItzB1ack/CalculatorYL2/internal/models"
)

func main() {
	app := models.New()
	fmt.Println("Сервер запущен на порту 8080")
	app.RunServer()
}
