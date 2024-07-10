package main

import "fmt"

func DortIslem(sayi1 int, sayi2 int) (int, int, int, int) {
	toplam := sayi1 + sayi2
	çikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolüm := sayi1 / sayi2
	return toplam, çikarim, carpim, bolüm
}
func main() {
	var sonuc1, sonuc2, sonuc3, sonuc4 = DortIslem(6, 2)
	// Yukarıda istemedğiniz bir işlem varsa onu sildiğimizde go dili bize hata verecektir o yüzden onun yerine " _ " kullanılır.
	fmt.Println("Toplam : ", sonuc1)
	fmt.Println("Cikarim : ", sonuc2)
	fmt.Println("Carpim : ", sonuc3)
	fmt.Println("Bolüm : ", sonuc4)
}
