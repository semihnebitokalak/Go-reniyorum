// Bu dersimizde yine güzel bir soru çözeceğiz...
/*
İki sayı birbirinin kendisi hariç bölenleri toplamına eşitse bu sayılara arkadaş sayılar denir.
Örnek: 220 ve 284
220 : 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 = 284 284 : 1 + 2 + 4 + 71 + 142 = 220 1636 yılında
Fermat 17296 ve 18416 nın arkadaş sayı olduklarını keşfetti.
Üçüncü çifti Descartes keşfetti.Leonhard Euler ise, kendi bulduğu 59 çift ile birlikte 62 çiftten oluşan bir liste hazırladı.
1866'da 16 yaşındaki İtalyan Nicolo Paganini 1184 ve 1210 sayılarının da böyle bir çift olduğunu gösterdi.
*/

package main

import "fmt"

func main() {
	// Bu soru için öncelikle elimizde 2 tane sayı olmalıdır.
	// Hemen bunları kullanıcıdan isteyelim.
	// Ayrıca tam sayı bölenlerinin toplamını tutmak için ayrı bir değişken atıyoruz.
	var sum1 int = 0
	var sum2 int = 0
	var sayi1 int
	var sayi2 int
	fmt.Print("Lütfen birinci sayiyi giriniz.")
	fmt.Scanln(&sayi1)
	fmt.Print("Lütfen ikinci sayiyi giriniz.")
	fmt.Scanln(&sayi2)
	// Şimdi yapacağımız şey ise girilen sayıların bölenleri bulmak olacaktır.
	// Bunun için bir for döngüsü açıp if ile kontrolleri sağlayacağız.
	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			sum1 = sum1 + i
		}
	}
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			sum2 = sum2 + i
		}
	}
	if sayi1 == sum2 && sayi2 == sum1 {
		fmt.Println("Girdiğiniz sayilar kardeş sayilardir.")
	} else {
		fmt.Println("Girilen sayilar kardeş sayi değildir.")
	}
}
