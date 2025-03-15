module github.com/ItzB1ack/CalculatorYL2

go 1.23.3

require (
	github.com/ItzB1ack/CalculatorYL2/internal v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require github.com/ItzB1ack/CalculatorYL2/pkg v0.0.0-20250224130152-20ff1c91bbb8 // indirect

replace github.com/ItzB1ack/CalculatorYL2/internal => ./internal

replace github.com/ItzB1ack/CalculatorYL2/pkg => ./pkg

replace github.com/ItzB1ack/CalculatorYL2/internal/models => ./internal/models
