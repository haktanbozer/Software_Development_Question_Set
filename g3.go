package main

// findMostRepeated, verilen bir string array içinde en sık tekrarlanan öğeyi bulan bir fonksiyondur.
func findMostRepeated(arr []string) string {
	// Öğelerin sayısını tutacak bir map tanımlanır.
	counts := make(map[string]int)

	// Array üzerinde dolaşılır ve her öğenin sayısı counts map'ine eklenir.
	for _, item := range arr {
		counts[item]++
	}

	// En sık tekrarlanan öğeyi ve sayısını tutacak değişkenler tanımlanır.
	var mostRepeated string
	var maxCount int

	// counts map'i üzerinde dolaşılır ve en sık tekrarlanan öğe bulunur.
	for item, count := range counts {
		if count > maxCount {
			mostRepeated = item
			maxCount = count
		}
	}

	return mostRepeated
}
