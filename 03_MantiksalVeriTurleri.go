// Bu dersimizde mantıksal türleri göreceğiz...
// Bir durum ya doğrudur ya da yanlıştır...
// Bu tür veriler boolean diye adlandırılır ve bool diye tanımlanır...

package main

import "fmt"

func main() {
	var durum bool
	var metin1 = "Engin"
	var metin2 = "Ahmet"
	// Buradaki iki eşittir mantıksal operatördür ve iki durumunun birbirine eşit olup olmamasını kontrol eder.
	// Mantıksal operatörlerde != ifadesi de yer alır ve iki değer eşit değilse anlamı katar.
	durum = metin1 == metin2

	fmt.Println(durum)
	// Eğer ünlem koyarsak aşağıdaki kod bloğundaki gibi verilen değerin tersini yazdıracaktır.
	fmt.Println(!durum)
	// Ayrıca alt satıra inmek için \n kullanabiliriz...

}
