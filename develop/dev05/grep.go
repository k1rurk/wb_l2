package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Regex interface {
	isExist(string) bool
}

type RawStr struct {
	str string
}

func (r *RawStr) isExist(strInFind string) bool {
	return strings.Contains(strInFind, r.str)
}

type RegexStr struct {
	str string
}

func (r *RegexStr) isExist(strInFind string) bool {
	res, err := regexp.MatchString(r.str, strInFind)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}

type Data struct {
	strMatch Regex
	file     *os.File
	flags    *SFlags
}

func NewData(strMatch Regex, file *os.File, flags *SFlags) *Data {
	return &Data{
		strMatch: strMatch,
		file:     file,
		flags:    flags,
	}
}

type SFlags struct {
	after,
	before,
	context int
	count,
	ignoreCase,
	invert,
	fixed,
	lineNum bool
}

func NewSFlags(after, before, context *int, count, ignoreCase, invert, fixed, lineNum *bool) *SFlags {
	return &SFlags{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
	}
}

type Line struct {
	line    string
	numLine int
}

func doBuffer(buff *[]Line, line Line) {
	if cap(*buff) == 0 {
		return
	}

	if len(*buff) < cap(*buff) {
		*buff = append(*buff, line)
	} else {
		*buff = append((*buff)[1:], line)
	}
}

func (d *Data) execute() {

	if d.flags.context != 0 {
		d.flags.after = d.flags.context
		d.flags.before = d.flags.context
	}

	seenLines := make(map[int]bool)

	lineNumber := 0
	countAfter := 0

	scanner := bufio.NewScanner(d.file)
	match := 0

	beforeBuffer := make([]Line, 0, d.flags.before)
	var linesForPrint []Line
	for scanner.Scan() {
		lineNumber++

		text := scanner.Text()
		exists := d.strMatch.isExist(strings.ToLower(text))

		if exists != d.flags.invert {
			match++
			linesForPrint = append(linesForPrint, beforeBuffer...)
			linesForPrint = append(linesForPrint, Line{line: text, numLine: lineNumber})
			beforeBuffer = beforeBuffer[:cap(beforeBuffer)]
			countAfter = d.flags.after
		} else {

			if countAfter != 0 {
				linesForPrint = append(linesForPrint, Line{line: text, numLine: lineNumber})
				countAfter--
			}

			doBuffer(&beforeBuffer, Line{line: text, numLine: lineNumber})
		}

		if !d.flags.count {
			for _, v := range linesForPrint {
				if _, ok := seenLines[v.numLine]; !ok {
					if d.flags.lineNum {
						fmt.Printf("%d:%v\n", v.numLine, v.line)
					} else {
						fmt.Printf("%v\n", v.line)
					}
					seenLines[v.numLine] = true
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if d.flags.count {
		fmt.Println("Count match lines:", match)
	}

}
