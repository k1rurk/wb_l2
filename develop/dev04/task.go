package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type char []rune

func (c char) Len() int {
	return len(c)
}

func (c char) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c char) Less(i, j int) bool {
	return c[i] < c[j]
}

func toLowerArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}
}

func execute(arr []string) map[string][]string {
	result := make(map[string][]string)
	//в словаре или дубликат
	checked := make(map[string]bool)

	toLowerArray(arr)

	for i, val := range arr {
		if _, ok := checked[val]; !ok {
			checked[val] = true
			for j := i + 1; j < len(arr); j++ {
				// проверка на дубликат
				if _, ok := checked[arr[j]]; !ok {
					if CheckTwoStrings(val, arr[j]) {
						checked[arr[j]] = true
						result[val] = append(result[val], arr[j])
					}
				}
			}
		}
	}

	for key, _ := range result {
		if len(result[key]) <= 1 {
			delete(result, key)
		}
		sort.SliceStable(result[key], func(i, j int) bool {
			return result[key][i] < result[key][j]
		})
	}
	return result
}

func CheckTwoStrings(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	strChars1 := make(char, len(str1))
	strChars2 := make(char, len(str2))

	copy(strChars1, char(str1))
	copy(strChars2, char(str2))

	sort.Sort(strChars1)
	sort.Sort(strChars2)

	for i := 0; i < len(strChars1); i++ {
		if strChars1[i] != strChars2[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []string{"пятак", "листок", "столик", "пятка", "слиток", "тяпка"}
	result := execute(arr)
	for key, val := range result {
		fmt.Printf("key=%v, slice=%v\n", key, val)
	}
}
