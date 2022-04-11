package main

import "testing"

var tests = []struct {
	str  string
	want string
}{
	{
		"a4bc2d5e", "aaaabccddddde",
	},
	{
		"abcd", "abcd",
	},
	{
		"45", "",
	},
	{
		"", "",
	},
	{
		"qwe\\4\\5", "qwe45",
	},
	{
		"qwe\\45", "qwe44444",
	},
	{
		"qwe\\\\5", "qwe\\\\\\\\\\",
	},
}

func TestStrings(t *testing.T) {
	for _, tstr := range tests {
		t.Run(tstr.str, func(t *testing.T) {
			ans, _ := Unpack(tstr.str)
			if ans != tstr.want {
				t.Errorf("got %s, want %s", ans, tstr.want)
			}
		})
	}
}
