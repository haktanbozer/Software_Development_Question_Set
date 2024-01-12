package main

import "testing"

func TestFindMostRepeated(t *testing.T) {
	// Test verisi
	input := []string{"apple", "pie", "apple", "red", "red", "red"}

	// findMostRepeated fonksiyonunu test et
	result := findMostRepeated(input)

	// Beklenen çıktı
	expectedOutput := "red"

	// Beklenen ve alınan çıktıları karşılaştır
	if result != expectedOutput {
		t.Errorf("Beklenen: %s, Alınan: %s", expectedOutput, result)
	}
}
