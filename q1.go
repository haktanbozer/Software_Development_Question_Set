package main //package main ifadesi, bir Go uygulamasının başlangıç noktasını belirtir. Bu ifade, derleyiciye "Bu dosya bir uygulama başlangıç noktasıdır ve buradan başla." demek için kullanılır. Yani, bir Go programı çalıştığında, main fonksiyonu çağrılır ve programın çalışması bu noktadan başlar.

import ( //Go dilinde import ifadesi, harici kütüphaneleri (paketleri) programa eklemek için kullanılır.

	"fmt"  //fmt kütüphanesi: Standart çıkışa yazma (printf vb.) ve diğer formatlama işlemleri için fonksiyonlar içerir.
	"sort" //sort kütüphanesi: Sıralama işlemleri için fonksiyonlar içerir.
)

func sortByA(words []string) []string { // Verilen  kelime listesini 'a' harfi sayısına göre azalan sırayla ve 'a' harfi sayısı eşitse uzunluğa göre artan sırayla sıralayan bir fonksiyon oluşturduk.

	sort.Slice(words, func(i, j int) bool {
		countA_i := countA(words[i])
		countA_j := countA(words[j])

		if countA_i == countA_j {
			return len(words[i]) < len(words[j])
		}

		return countA_i > countA_j
	})

	return words
}

func countA(word string) int { // Bir kelimenin içindeki 'a' harfi sayısını hesaplayan bir fonksiyon oluşturduk.
	count := 0
	for _, char := range word {
		if char == 'a' {
			count++
		}
	}
	return count
}

func main() { // Programın başlangıç noktası.

	// Örnek giriş verisi
	inputWords := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// sortByA fonksiyonunu kullanarak giriş verisini sırala
	outputWords := sortByA(inputWords)

	// Sonucu ekrana yazdır
	fmt.Println(outputWords)
}
