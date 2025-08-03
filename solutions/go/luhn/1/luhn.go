package luhn

import (
    "strconv"
    "strings"
    "unicode"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
        return false
    }

    sum := 0
    double := false

    for i := len(id) - 1; i >= 0; i-- {
        num := rune(id[i])
        if !unicode.IsDigit(num) {
            return false
        }

        digit, _ := strconv.Atoi(string(num))

        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

        sum += digit
        double = !double
    }

    return sum%10 == 0
}
