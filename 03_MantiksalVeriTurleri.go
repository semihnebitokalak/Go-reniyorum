// Bu dersimizde mantıksal türleri göreceğiz...
// Bir durum ya doğrudur ya da yanlıştır...
// Bu tür veriler boolean diye adlandırılır ve bool diye tanımlanır...

package main

import "fmt"

func main() {
	var durum1 bool
	var durum2 bool
	var metin1 = "Engin"
	var metin2 = "Ahmet"
	var metin3 = "Engin"
	// Buradaki iki eşittir mantıksal operatördür ve iki durumunun birbirine eşit olup olmamasını kontrol eder.
	// Mantıksal operatörlerde != ifadesi de yer alır ve iki değer eşit değilse anlamı katar.
	durum1 = metin1 == metin2
	durum2 = metin1 == metin3

	fmt.Println(durum1)
	// Eğer ünlem koyarsak aşağıdaki kod bloğundaki gibi verilen değerin tersini yazdıracaktır.
	fmt.Println(!durum1)
	// Ayrıca alt satıra inmek için \n kullanabiliriz...
	fmt.Println(durum2)

}
