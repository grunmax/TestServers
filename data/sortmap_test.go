package data

import "testing"

func Test_getTopByValue(t *testing.T) {
	testieMap := map[string]int{
		"bird":  4,
		"snake": 1,
		"cat":   2,
		"dog":   4,
	}
	okData := []string{"bird", "dog", "cat"}
	someData := getTopByValue(testieMap, 3)
	for i, v := range someData {
		if v != okData[i] {
			t.Error("Test: top words")
		}
	}
}
