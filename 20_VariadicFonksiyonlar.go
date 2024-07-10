// Go dilinde variadic fonksiyon, değişken sayıda argüman alabilen fonksiyonlardır.
// Bir variadic fonksiyon tanımlamak için fonksiyon parametrelerinden birinin türünden önce üç nokta (...) eklenir.
// Bu parametre, fonksiyona bir slice (dilim) olarak geçer.
package main

import "fmt"

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	return toplam

}
func main() {
	sonuc1 := ToplaVariadic(1, 2, 3, 4, 5)
	sonuc2 := ToplaVariadic(2, 4)
	sonuc3 := ToplaVariadic()
	fmt.Println("Toplam : ", sonuc1)
	fmt.Println("Toplam : ", sonuc2)
	fmt.Println("Toplam : ", sonuc3)
	// Eğer dizi tanımlayıp onu toplam olarak yazdırırsak ufak bi trick var hemen ona bakalım.
	sayilar := []int{1, 2, 4, 56, 2, 0, 4}
	// Fonsiyona gönderme yaparken dizilerde sonuna "..." nokta konur
	fmt.Println("Toplam :", ToplaVariadic(sayilar...))

}
