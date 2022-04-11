package main

import "testing"

func TestAnagram(t *testing.T) {
	c1 := "столик"
	cS1 := []string{"листок", "слиток", "столки"}
	for _, val := range cS1 {
		if !CheckTwoStrings(val, c1) {
			t.Errorf("Test #1 Faild")
		}
	}

	c2 := "клоун"
	cS2 := []string{"кулон", "лукно", "уклон"}
	for _, val := range cS2 {
		if !CheckTwoStrings(val, c2) {
			t.Errorf("Test #2 Faild")
		}
	}

	c3 := "карта"
	cS3 := []string{"карат", "нарок", "катар"}
	isOdd := false
	for _, val := range cS3 {
		if !CheckTwoStrings(val, c3) {
			isOdd = true
			break
		}
	}
	if !isOdd {
		t.Errorf("Test #3 Faild")
	}
}
