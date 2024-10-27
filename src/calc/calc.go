package calc

func Calc(expression string) (float64, error) {
    tokens, err := parse(expression)
    if err != nil {
        return 0, err
    }
    result, err := evaluate(tokens)
    if err != nil {
        return 0, err
    }
    return result, nil
}