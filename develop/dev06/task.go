package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type fieldsList []int

func (f *fieldsList) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *fieldsList) Set(value string) error {
	spltArr := strings.Split(value, ",")
	for _, val := range spltArr {
		tmp, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		*f = append(*f, tmp)
	}
	return nil
}

type SFlags struct {
	fields    fieldsList
	delimiter string
	separated bool
}

func NewSFlags(fields fieldsList, delimiter *string, separated *bool) *SFlags {
	return &SFlags{
		fields:    fields,
		delimiter: *delimiter,
		separated: *separated,
	}
}

func (f *SFlags) chooseFields(strArr []string) []string {
	if len(f.fields) == 0 {
		return strArr
	}
	var res []string
	l := len(strArr)
	for _, val := range f.fields {
		if l < val {
			break
		}
		res = append(res, strArr[val-1])
	}
	return res
}

func (f *SFlags) chooseDelimiter(str string) []string {
	sArr := strings.Split(str, f.delimiter)
	return sArr
}

func (f *SFlags) isSeparated(arr []string) bool {
	return len(arr) > 1
}

func execute(strArr []string, flg *SFlags) [][]string {
	var result [][]string
	var tempArrStr []string
	for _, str := range strArr {
		tempArrStr = flg.chooseDelimiter(str)
		if flg.separated && !flg.isSeparated(tempArrStr) {
			continue
		}
		tempArrStr = flg.chooseFields(tempArrStr)
		result = append(result, tempArrStr)
	}

	return result
}

func main() {
	var fieldsPtr fieldsList
	flag.Var(&fieldsPtr, "f", "выбрать поля (колонки).")
	delimiterPtr := flag.String("d", "\t", "использовать другой разделитель.")
	separatedPtr := flag.Bool("s", false, "только строки с разделителем.")
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Println("usage: [flags] string")
		flag.PrintDefaults()
		os.Exit(-1)
	}
	inStrings := strings.Split(args[0], "\n")
	flg := NewSFlags(fieldsPtr, delimiterPtr, separatedPtr)
	res := execute(inStrings, flg)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%v\n", res[i])
	}
}
