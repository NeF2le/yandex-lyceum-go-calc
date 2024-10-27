package calc

import (
    "errors"
    "strconv"
)

// Функция для задания приоритета операций
func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    }
    return 0
}

// Функция для выполнения арифметических операций
func applyOp(a, b float64, op string) (float64, error) {
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    }
    return 0, errors.New("invalid operator")
}

// Основная функция Evaluate, выполняющая вычисление
func evaluate(tokens []string) (float64, error) {
    var values []float64        // Стек для чисел
    var ops []string            // Стек для операторов

    for _, token := range tokens {
        // Если токен - число
        if num, err := strconv.ParseFloat(token, 64); err == nil {
            values = append(values, num)
        } else if token == "(" {
            ops = append(ops, token)
        } else if token == ")" {
            // Вычисляем до открытия скобки
            for len(ops) > 0 && ops[len(ops)-1] != "(" {
                if len(values) < 2 {
                    return 0, errors.New("invalid expression")
                }
                b, a := values[len(values)-1], values[len(values)-2]
                op := ops[len(ops)-1]
                values = values[:len(values)-2]
                ops = ops[:len(ops)-1]
                result, err := applyOp(a, b, op)
                if err != nil {
                    return 0, err
                }
                values = append(values, result)
            }
            // Удаляем "(" из стека операторов
            if len(ops) == 0 || ops[len(ops)-1] != "(" {
                return 0, errors.New("mismatched parentheses")
            }
            ops = ops[:len(ops)-1]
        } else {
            // Токен - оператор
            for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(token) {
                if len(values) < 2 {
                    return 0, errors.New("invalid expression")
                }
                b, a := values[len(values)-1], values[len(values)-2]
                op := ops[len(ops)-1]
                values = values[:len(values)-2]
                ops = ops[:len(ops)-1]
                result, err := applyOp(a, b, op)
                if err != nil {
                    return 0, err
                }
                values = append(values, result)
            }
            ops = append(ops, token)
        }
    }

    // Выполняем оставшиеся операции в стеке
    for len(ops) > 0 {
        if len(values) < 2 {
            return 0, errors.New("invalid expression")
        }
        b, a := values[len(values)-1], values[len(values)-2]
        op := ops[len(ops)-1]
        values = values[:len(values)-2]
        ops = ops[:len(ops)-1]
        result, err := applyOp(a, b, op)
        if err != nil {
            return 0, err
        }
        values = append(values, result)
    }

    // Результат - последний элемент в стеке чисел
    if len(values) != 1 {
        return 0, errors.New("invalid expression")
    }
    return values[0], nil
}
