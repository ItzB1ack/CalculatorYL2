module github.com/ItzB1ack/CalculatorYL2

go 1.23.3

require (
	github.com/ItzB1ack/CalculatorYL2/internal v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

replace github.com/ItzB1ack/CalculatorYL2/internal => ./internal

replace github.com/ItzB1ack/CalculatorYL2/pkg => ./pkg
