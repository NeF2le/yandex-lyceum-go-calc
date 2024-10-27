package calc

import (
    "errors"
    "unicode"
	"strings"
)

func parse(expression string) ([]string, error) {
    var tokens []string
    var number string
	expression = strings.Replace(expression, " ", "", -1)

    for i, char := range expression {
        switch {
        case unicode.IsDigit(char) || char == '.':
            number += string(char)

        case char == '+' || char == '*' || char == '/' || char == '(' || char == ')':
            if number != "" {
                tokens = append(tokens, number)
                number = ""
            }
            tokens = append(tokens, string(char))

		case char == '-':
            if number != "" {
                tokens = append(tokens, number)
                number = ""
            }

            if i == 0 || expression[i-1] == '(' {
                number = "-"
            } else {
                tokens = append(tokens, string(char))
            }

        default:
            return nil, errors.New("invalid character in expression")
        }
    }
    if number != "" {
        tokens = append(tokens, number)
    }
    return tokens, nil
}