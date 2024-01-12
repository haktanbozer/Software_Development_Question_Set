package main

import "fmt"

func generateOutput(n int) { // recursive olarak çağrılarak belirtilen formatta çıktı üreten bir fonksiyon oluşturdum.

	if n > 0 { // Eğer n sıfırdan büyükse devam et.

		generateOutput(n / 2) //  n'yi 2'ye bölerek tekrar çağır.

		fmt.Println(n) // Her seviyede elde edilen n değerini ekrana yazdır.
	}
}

func main() {
	// Örnek giriş
	input := 9

	generateOutput(input) // generateOutput fonksiyonunu örnek girişle çağır.
}
