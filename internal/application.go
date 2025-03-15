package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	calc "github.com/ItzB1ack/CalculatorYL2/pkg"
)

type Config struct {
	Address string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Address = os.Getenv("PORT")

	if config.Address == "" {
		config.Address = "8080"
	}

	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		request := new(Request)
		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Запрос невалиден"}`)
			return
		}

		result, err := calc.Calc(request.Expression)

		if err != nil {
			errMsg := err.Error()
			if strings.Contains(errMsg, "brackets") || strings.Contains(errMsg, "expression") || strings.Contains(errMsg, "zero") {
				w.WriteHeader(http.StatusUnprocessableEntity)
				response := Response{Error: errMsg}
				json.NewEncoder(w).Encode(response)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				response := Response{Error: "Внутренняя ошибка сервера"}
				json.NewEncoder(w).Encode(response)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		response := Response{Result: result}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := Response{Error: "Данный метод не поддерживается"}
		json.NewEncoder(w).Encode(response)
	}

}

func (a *Application) StopServer() error {
	http.HandleFunc("/", CalcHandler)
	return nil
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Address, nil)
}
