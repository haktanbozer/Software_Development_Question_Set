package main

import (
	"reflect"
	"testing"
)

func TestSortByA(t *testing.T) {
	inputWords := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	expectedOutput := []string{"l", "a", "aab", "ef", "fdz", "kf", "zc", "aaabcd", "aaaasd", "cssssssd", "lklklklklklklklkl"}

	outputWords := sortByA(inputWords)

	if !reflect.DeepEqual(outputWords, expectedOutput) {
		t.Errorf("Beklenen: %v, Alınan: %v", expectedOutput, outputWords)
	}
}
