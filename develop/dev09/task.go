package main

import (
	"log"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Println("usage: site")
		os.Exit(-1)
	}

	strUrl := args[0]

	getSite(strUrl)
}
