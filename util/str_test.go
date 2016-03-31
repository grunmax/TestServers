package util

import (
	"fmt"
	"testing"
)

func TestRegSplit(t *testing.T) {
	testieData := "alfa bravo    tango"
	list := RegSplit(testieData, "[^\\S]+")
	if len(list) != 3 {
		t.Error(fmt.Sprintf("Test: can`t split: %s", testieData))
	}
}

func TestWordsCheckList(t *testing.T) {
	testieData := []string{"alfa", "танго", "sa"}
	listOk, listBad := WordsCheckList(testieData, 4)
	if !(len(listOk) == 2) && (len(listBad) == 1) {
		t.Error(fmt.Sprintf("Test: wrong filter words: %v", testieData))
	}
}
