package main

import (
	"fmt"
	"github.com/ItzB1ack/CalculatorYL/internal"
)

func main() {
	app := internal.New()
	fmt.Println("Сервер запущен на порту 8080")
	app.RunServer()
}
