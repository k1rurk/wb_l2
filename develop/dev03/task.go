package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	columnPtr := flag.Int("k", 1, "указание колонки для сортировки (слова в строке могут\nвыступать в качестве колонок, по умолчанию разделитель —\nпробел).")
	numericPtr := flag.Bool("n", false, "сортировать по числовому значению.")
	reversePtr := flag.Bool("r", false, "сортировать в обратном порядке.")
	repeatPtr := flag.Bool("u", false, "не выводить повторяющиеся строки.")
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Println("usage: [flags] file")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	file, err := ioutil.ReadFile(args[0])

	if err != nil {
		log.Fatalln(err)
	}

	strSplit := strings.Split(string(file), "\n")

	for i := range strSplit {
		strSplit[i] = strings.TrimSuffix(strSplit[i], "\r")
	}

	flags := &Flags{
		column:  *columnPtr,
		numeric: *numericPtr,
		reverse: *reversePtr,
		repeat:  *repeatPtr,
	}

	arr, err := flags.sortStrings(strSplit)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v1 := range arr {
		fmt.Printf("%v\n", v1)
	}
}
