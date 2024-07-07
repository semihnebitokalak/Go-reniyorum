// Koşul ve şart örnek soru çözümü!
// Şimdi beraber deneme yapacağız...
// Seneryomuz şu 3 adet int değişken tanımlayacağız...
// Ekrana en büyük olanı yazdıracağız...

package main

import "fmt"

func main() {
	// Tanımlamayı şu şekilde de yapabilirsiniz...
	// var sayi1, sayi2, sayi3 int = 12, 42, 25
	sayi1 := 25
	sayi2 := 53
	sayi3 := 30
	if sayi1 > sayi2 {
		if sayi1 > sayi3 {
			fmt.Println("En büyük sayi sayi1 dir ve değeri = " + fmt.Sprintf("%v", sayi1))
		}
		fmt.Println("En büyük sayi sayi3 tür ve değeri = " + fmt.Sprintf("%v", sayi3))
	}
	if sayi2 > sayi3 {
		fmt.Println("En büyük sayi sayi2 dir ve değeri = " + fmt.Sprintf("%v", sayi2))
	}
}

/*
Siz bu kodu sayi1 için iki tane if yapısı, sayi2 için iki tane if,
sayi3 için iki tane if yapısı oluşturarak uzun uzadıya yapabilirsiniz.
Veya sayi1'e en büyük diye bi değer atayıp diğer değişkenlerin sayi1'den
büyük olup olmadığını kontrol edebilirsiniz.
*/
