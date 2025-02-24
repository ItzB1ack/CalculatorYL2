## HTTP-сервис калькулятор

Веб-сервис на Go для вычисления математических выражений. Поддерживает базовые арифметические операции: сложение, вычитание, умножение, деление и работу со скобками.

## Описание задачи

Проект представляет собой распределенный вычислитель арифметических выражений. Пользователь вводит строку с выражением (например, `2 + 2 * 2`), и система возвращает результат с учетом приоритетов операций.

### Поддерживаемые операции
- Сложение (`+`)
- Вычитание (`-`)
- Умножение (`*`)
- Деление (`/`)

### Пример использования
```
Введите выражение: 2 + 2 * 2
Результат: 6
```

## Структура проекта

- `cmd/` - директория с файлом main.go
- `internal/application/` - директория с кодом сервера и тестами для проверки работы сервера
- `pkg/calculation/` - директория с кодом калькулятора и тестами для проверки работы калькулятора

## Требования

- Go 1.20 или выше
- Git

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/ItzB1ack/CalculatorYL2.git -b master
cd Calculator_WebApp
```

2. Соберите проект:
```bash
go build -o calculator ./cmd/main.go
```

3. Запустите сервер:
```bash
./calculator
```

Сервер по умолчанию запускается на порту 8080.

## Работа с сервисом

### API Документация

Сервис предоставляет единственную конечную точку:

**POST /calculate**

Принимает JSON-запрос с математическим выражением и возвращает результат вычисления.

#### Формат запроса:
```json
{
    "expression": "строка с математическим выражением"
}
```

#### Формат ответа:
```json
{
    "result": число,
    "error": "сообщение об ошибке (опционально)"
}
```

### Поддерживаемые операции

- Сложение (+)
- Вычитание (-)
- Умножение (*)
- Деление (/)
- Скобки ( )

### Ограничения

- Поддерживаются только целые числа
- Максимальная длина выражения: 100 символов
- Максимальная глубина вложенности скобок: 10

## Примеры работы с сервисом

### Корректный запрос:
```json
POST /calculate
{
    "expression": "1+2*3"
}

Ответ:
{
    "result": 7
}
```

### Запросы с невалидным выражением

1. Ошибка в скобках:
```json
{
    "expression": "1+((2+3+4)*5"
}

Ответ:
{
    "error": "Некорректное количество скобок"
}
```

2. Ошибка в выражении:
```json
{
    "expression": "1+2+3+-4*5"
}

Ответ:
{
    "error": "Некорректное выражение"
}
```

3. Деление на ноль:
```json
{
    "expression": "(2+3)/0"
}

Ответ:
{
    "error": "Деление на ноль"
}
```

## Тестирование

Для запуска всех тестов выполните команду:
```bash
go test ./...
```

## Запуск проекта
1. Убедитесь, что установлен Go.
2. Клонируйте репозиторий.
3. Перейдите в директорию проекта и выполните команду:
```
go run cmd/main.go
