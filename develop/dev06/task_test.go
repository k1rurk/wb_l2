package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type testpair struct {
	input  string
	flags  SFlags
	answer [][]string
}

var tests = []testpair{
	{
		"Таинственная река (2003)\nДевушка в тумане (2017)\nТело (2012)\nБагровые реки (2000)",
		SFlags{
			fields:    fieldsList{1, 3},
			delimiter: " ",
		},
		[][]string{
			{"Таинственная", "(2003)"},
			{"Девушка", "тумане"},
			{"Тело"},
			{"Багровые", "(2000)"},
		},
	},
}

func TestAverage(t *testing.T) {
	for num, pair := range tests {
		strSplt := strings.Split(pair.input, "\n")
		res := execute(strSplt, &pair.flags)
		for i := range res {
			if !reflect.DeepEqual(res[i], pair.answer[i]) {
				t.Error(
					fmt.Sprintf("Test#%d", num+1), "\n",
					"expected", pair.answer[i], "\n",
					"got", res[i],
				)
			}
		}
	}
}
