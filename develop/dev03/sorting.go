package main

import (
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	numeric bool
	reverse bool
	column  int
	repeat  bool
}

func (f *Flags) sortStrings(arr []string) ([]string, error) {

	if f.repeat {
		arr = checkRepeatStrings(arr)
	}

	if f.column > 1 {
		sortByColumn(arr, f.column, f.numeric)
	} else {
		sortDefault(arr, f.numeric)
	}

	if f.reverse {
		reverseSlice(arr)
	}

	return arr, nil
}

func checkRepeatStrings(arr []string) []string {
	var res []string
	set := make(map[string]bool)
	for i := range arr {
		set[arr[i]] = true
	}
	for key := range set {
		res = append(res, key)
	}
	return res
}

func sortByColumn(arr []string, column int, numeric bool) {
	arrSplitWords := make([][]string, len(arr))
	for i := range arr {
		arrSplitWords[i] = strings.Fields(arr[i])
	}
	var col1, col2 int
	var lenArr1, lenArr2 int
	sort.SliceStable(arrSplitWords, func(i, j int) bool {
		lenArr1, lenArr2 = len(arrSplitWords[i]), len(arrSplitWords[j])
		if lenArr1 <= 1 {
			col1 = 0
		} else {
			if column > lenArr1 {
				col1 = lenArr1 - 1
			} else {
				col1 = column - 1
			}
		}
		if lenArr2 <= 1 {
			col2 = 0
		} else {
			if column > lenArr2 {
				col2 = lenArr2 - 1
			} else {
				col2 = column - 1
			}
		}
		if numeric {
			if num1, err := strconv.Atoi(arrSplitWords[i][col1]); err == nil {
				if num2, err := strconv.Atoi(arrSplitWords[j][col2]); err == nil {
					return num1 < num2
				}
			}
			return arrSplitWords[i][col1] < arrSplitWords[j][col2]
		} else {
			return arrSplitWords[i][col1] < arrSplitWords[j][col2]
		}
	})
	for i := range arr {
		arr[i] = strings.Join(arrSplitWords[i], " ")
	}
}

func sortDefault(arr []string, numeric bool) {
	arrSplitWords := make([][]string, len(arr))
	for i := range arr {
		arrSplitWords[i] = strings.Fields(arr[i])
	}
	sort.SliceStable(arrSplitWords, func(i, j int) bool {
		if numeric {
			if num1, err := strconv.Atoi(arrSplitWords[i][0]); err == nil {
				if num2, err := strconv.Atoi(arrSplitWords[j][0]); err == nil {
					return num1 < num2
				}
			}
			return arrSplitWords[i][0] < arrSplitWords[j][0]
		} else {
			return arrSplitWords[i][0] < arrSplitWords[j][0]
		}
	})
	for i := range arr {
		arr[i] = strings.Join(arrSplitWords[i], " ")
	}
}

func reverseSlice(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
