package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(str string) (string, error) {

	if len(str) == 0 {
		return "", nil
	}
	builder := strings.Builder{}

	esc := false
	escCount := 0

	for i, val := range str {
		if unicode.IsDigit(val) && i == 0 {
			return "", fmt.Errorf(str, " неккоректная строка")
		}

		if unicode.IsDigit(val) {
			digit, _ := strconv.Atoi(string(val))
			if esc && escCount <= 1 {
				builder.WriteRune(val)
				escCount++
			} else {
				for j := 0; j < digit-1; j++ {
					builder.WriteRune(rune(str[i-1]))
				}
				escCount = 0
				esc = false
			}
		} else if 92 == val {
			if str[i-1] != 92 {
				esc = true
				if escCount <= 2 {
					escCount = 1
				}
			} else {
				escCount++
				builder.WriteRune(val)
			}
		} else {
			builder.WriteRune(val)
		}
	}

	return builder.String(), nil
}
