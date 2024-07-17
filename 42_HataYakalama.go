package main

import (
	"fmt"
	"os"
)

func DosyaDurumu() {
	// Aşağıdaki kod bloğu nedir ne ifade eder onu inceleyelim.
	// os.Open ile dosya oluşturduk.
	// Eğer dosya bulunmazsa err(error) olarak hata vericek. Error dönücek.
	// Dosya varsa dosyayı veriyor.
	f, err := os.Open("demo1.txt")
	// Dosyayı bulunursa err = nill (error ın o andaki haline nill denir)
	if err != nil {
		fmt.Println("Dosya bulunamadi.")
	} else {
		fmt.Println(f.Name())
	}
}
func DosyaOluşturma() {
	// Dosya oluşturma
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Dosya oluşturulamadi:", err)
	} else {
		fmt.Println("Dosya oluşturuldu:", file.Name())
	}
}
func main() {
	DosyaDurumu()
	DosyaOluşturma()
}
