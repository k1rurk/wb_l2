package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	afterPtr := flag.Int("A", 0, "печатать +N строк после совпадения.")
	beforePtr := flag.Int("B", 0, "печатать +N строк до совпадения.")
	contextPtr := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения.")
	countPtr := flag.Bool("c", false, "количество строк.")
	ignoreCasePtr := flag.Bool("i", false, "игнорировать регистр.")
	invertPtr := flag.Bool("v", false, "вместо совпадения, исключать.")
	fixedPtr := flag.Bool("F", false, "точное совпадение со строкой, не паттерн.")
	lineNumPtr := flag.Bool("n", false, "печатать номер строки.")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Println("usage: [flags] string")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	flags := NewSFlags(afterPtr, beforePtr, contextPtr,
		countPtr, ignoreCasePtr, invertPtr, fixedPtr, lineNumPtr)

	filename := args[1]
	fInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	if fInfo.Size() == 0 {
		log.Fatalln("file is empty")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	rMatch := args[0]
	var regMatch Regex
	if flags.fixed {
		if flags.ignoreCase {
			rMatch = "(?i)" + rMatch
		}
		regMatch = &RawStr{str: rMatch}
	} else {
		if flags.ignoreCase {
			rMatch = strings.ToLower(rMatch)
		}
		regMatch = &RegexStr{str: rMatch}
	}

	data := NewData(regMatch, file, flags)

	data.execute()

}
