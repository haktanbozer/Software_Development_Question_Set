package main

import (
	"io/ioutil"
	"os"
	"testing"
)

//goland:noinspection ALL
func TestGenerateOutput(t *testing.T) {
	// Orijinal Stdout'u sakla
	originalStdout := os.Stdout

	// Geçici bir dosya oluştur
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatalf("TempFile creation failed: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tmpfile.Name())

	// Stdout'u geçici olarak dosyaya yönlendir
	os.Stdout = tmpfile

	// Testin sonunda orijinal Stdout'u geri yükle ve dosyayı kapat
	defer func() {
		os.Stdout = originalStdout
		err := tmpfile.Close()
		if err != nil {
			return
		}
	}()

	// generateOutput fonksiyonunu test et
	generateOutput(9)

	// Dosyadaki içeriği oku
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Beklenen çıktı
	expectedOutput := "1\n2\n4\n8\n9\n"

	// Beklenen ve alınan çıktıları karşılaştır
	if string(content) != expectedOutput {
		t.Errorf("Beklenen: %s, Alınan: %s", expectedOutput, content)
	}
}
