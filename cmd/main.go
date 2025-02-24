package main

import (
	"fmt"

	application "github.com/ItzB1ack/CalculatorYL2/internal"
)

func main() {
	app := application.New()
	fmt.Println("Сервер запущен на порту 8080")
	app.RunServer()
}
