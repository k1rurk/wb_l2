package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testSort struct {
	strFile []string
	result  []string
}

var testColumn = []testSort{
	{
		strFile: []string{
			"ugabuga 21345 dfffg",
			"abucs 342 kk",
			"bbb",
			"dfs sdaf o",
			"000 adsf u",
		},
		result: []string{
			"ugabuga 21345 dfffg",
			"abucs 342 kk",
			"000 adsf u",
			"bbb",
			"dfs sdaf o",
		},
	},
}

var testNumeric = []testSort{
	{
		strFile: []string{
			"ugabuga 21345 dfffg",
			"abucs 342 kk",
			"bbb",
			"dfs sdaf o",
			"000 adsf u",
		},
		result: []string{
			"abucs 342 kk",
			"ugabuga 21345 dfffg",
			"000 adsf u",
			"bbb",
			"dfs sdaf o",
		},
	},
}

var testReverse = []testSort{
	{
		strFile: []string{
			"JANsdfnskdfnkj jurr",
			"FEBsakndqwndnkj JAss",
			"Alkdnwldnknk Fuukj",
			"asd Sors",
			"Michigan 21312",
			"asd",
			"asd",
			"ugabuga 505",
		},
		result: []string{
			"ugabuga 505",
			"asd",
			"asd",
			"asd Sors",
			"Michigan 21312",
			"JANsdfnskdfnkj jurr",
			"FEBsakndqwndnkj JAss",
			"Alkdnwldnknk Fuukj",
		},
	},
}

var testUnique = []testSort{
	{
		strFile: []string{
			"JANsdfnskdfnkj jurr",
			"FEBsakndqwndnkj JAss",
			"Alkdnwldnknk Fuukj",
			"asd Sors",
			"Michigan 21312",
			"asd",
			"asd",
			"ugabuga 505",
		},
		result: []string{
			"Alkdnwldnknk Fuukj",
			"FEBsakndqwndnkj JAss",
			"JANsdfnskdfnkj jurr",
			"Michigan 21312",
			"asd Sors",
			"asd",
			"ugabuga 505",
		},
	},
}

func TestColumns(t *testing.T) {
	for _, test := range testColumn {
		sortByColumn(test.strFile, 2, false)
		assert.Equal(t, test.strFile, test.result)
	}
}

func TestNumeric(t *testing.T) {
	for _, test := range testNumeric {
		sortByColumn(test.strFile, 2, true)
		assert.Equal(t, test.strFile, test.result)
	}
}

func TestReverse(t *testing.T) {
	for _, test := range testReverse {
		sortDefault(test.strFile, false)
		reverseSlice(test.strFile)
		assert.Equal(t, test.strFile, test.result)
	}
}

func TestUnique(t *testing.T) {
	for _, test := range testUnique {
		sortDefault(test.strFile, false)
		test.strFile = checkRepeatStrings(test.strFile)
		assert.Equal(t, test.strFile, test.result)
	}
}
